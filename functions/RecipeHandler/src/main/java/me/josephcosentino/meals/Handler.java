package me.josephcosentino.meals;

import com.amazonaws.services.lambda.runtime.Context;
import com.amazonaws.services.lambda.runtime.RequestHandler;
import com.amazonaws.services.lambda.runtime.events.APIGatewayV2ProxyRequestEvent;
import com.amazonaws.services.lambda.runtime.events.APIGatewayV2ProxyResponseEvent;
import lombok.SneakyThrows;
import lombok.extern.slf4j.Slf4j;
import me.josephcosentino.meals.handlers.ResourceHandler;
import me.josephcosentino.meals.handlers.ResourceHandlers;
import me.josephcosentino.meals.model.ErrorResponse;
import me.josephcosentino.meals.util.Environment;
import me.josephcosentino.meals.util.Jackson;
import software.amazon.awssdk.http.HttpStatusCode;
import software.amazon.awssdk.http.SdkHttpMethod;

import java.util.Map;

/**
 * Handler for requests to Lambda function.
 */
@Slf4j
public class Handler implements RequestHandler<APIGatewayV2ProxyRequestEvent, APIGatewayV2ProxyResponseEvent> {

    // TODO refactor. this way too complicated/convoluted
    private static final ResourceHandler<?> RESOURCE_HANDLER;

    static {
        final var resourceName = Environment.get("resource");
        final var supplier = ResourceHandlers.get(resourceName);
        if (supplier == null) {
            throw new IllegalStateException("No supported resource found for " + resourceName);
        }
        RESOURCE_HANDLER = supplier.get();
    }

    public APIGatewayV2ProxyResponseEvent handleRequest(APIGatewayV2ProxyRequestEvent input,
                                                        Context context) {
        try {
            return doHandleRequest(input, context);
        } catch (Exception e) {
            log.error("Handle failed", e);
            return internalServerError();
        }
    }

    private APIGatewayV2ProxyResponseEvent doHandleRequest(APIGatewayV2ProxyRequestEvent input,
                                                           Context context) {
        switch (SdkHttpMethod.fromValue(input.getHttpMethod())) {
            case GET:
                return get(input);
            case POST:
                return post(input);
            case PUT:
                return put(input);
            case DELETE:
                return delete(input);
            default:
                return badRequest();
        }
    }

    @SneakyThrows
    private <T> APIGatewayV2ProxyResponseEvent get(APIGatewayV2ProxyRequestEvent input) {
        final var resource = Handler.RESOURCE_HANDLER.getById(getIdFromPath(input)).get();
        log.info("GET: {}", resource);
        if (resource == null) {
            return notFound();
        }
        return success(resource);
    }

    @SneakyThrows
    @SuppressWarnings("unchecked")
    private <T> APIGatewayV2ProxyResponseEvent put(APIGatewayV2ProxyRequestEvent input) {
        final var bodyResource = ((ResourceHandler<T>) Handler.RESOURCE_HANDLER).asResource(input.getBody());
        var updatedResource = ((ResourceHandler<T>) Handler.RESOURCE_HANDLER).update(bodyResource).get();
        log.info("PUT: {}", updatedResource);
        return success(updatedResource);
    }

    @SneakyThrows
    @SuppressWarnings("unchecked")
    private <T> APIGatewayV2ProxyResponseEvent post(APIGatewayV2ProxyRequestEvent input) {
        final var bodyResource = ((ResourceHandler<T>) Handler.RESOURCE_HANDLER).asResource(input.getBody());
        var updatedResource = ((ResourceHandler<T>) Handler.RESOURCE_HANDLER).create(bodyResource).get();
        log.info("POST: {}", updatedResource);
        return success(updatedResource);
    }

    @SneakyThrows
    @SuppressWarnings("unchecked")
    private <T> APIGatewayV2ProxyResponseEvent delete(APIGatewayV2ProxyRequestEvent input) {
        final var resource = ((ResourceHandler<T>) Handler.RESOURCE_HANDLER).delete(getIdFromPath(input)).get();
        log.info("DELETE: {}", resource);
        if (resource == null) {
            return notFound();
        }
        return success(resource);
    }

    private String getIdFromPath(APIGatewayV2ProxyRequestEvent input) {
        return input.getPathParameters().get("id");
    }

    private APIGatewayV2ProxyResponseEvent success(Object content) {
        final var resp = new APIGatewayV2ProxyResponseEvent();
        resp.setBody(toJsonString(content));
        resp.setHeaders(Map.of("Content-Type", "application/json"));
        resp.setStatusCode(HttpStatusCode.OK);
        return resp;
    }

    private APIGatewayV2ProxyResponseEvent notFound() {
        final var resp = new APIGatewayV2ProxyResponseEvent();
        resp.setHeaders(Map.of("Content-Type", "application/json"));
        resp.setStatusCode(HttpStatusCode.NOT_FOUND);
        return resp;
    }

    private APIGatewayV2ProxyResponseEvent badRequest() {
        final var resp = new APIGatewayV2ProxyResponseEvent();
        resp.setBody(createErrorResponseBody("Bad Request"));
        resp.setHeaders(Map.of("Content-Type", "application/json"));
        resp.setStatusCode(HttpStatusCode.BAD_REQUEST);
        return resp;
    }

    private APIGatewayV2ProxyResponseEvent internalServerError() {
        final var resp = new APIGatewayV2ProxyResponseEvent();
        resp.setBody(createErrorResponseBody("Internal Server Error"));
        resp.setHeaders(Map.of("Content-Type", "application/json"));
        resp.setStatusCode(HttpStatusCode.INTERNAL_SERVER_ERROR);
        return resp;
    }

    private String createErrorResponseBody(String message) {
        try {
            return toJsonString(ErrorResponse.builder().message(message).build());
        } catch (Exception e) {
            log.error("Failed to write value as string", e);
            return "{}";
        }
    }

    @SneakyThrows
    private String toJsonString(Object obj) {
        return Jackson.getObjectMapper().writeValueAsString(obj);
    }
}
