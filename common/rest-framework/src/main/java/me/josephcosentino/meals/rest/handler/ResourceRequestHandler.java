package me.josephcosentino.meals.rest.handler;

import com.amazonaws.services.lambda.runtime.Context;
import com.amazonaws.services.lambda.runtime.RequestHandler;
import com.amazonaws.services.lambda.runtime.events.APIGatewayProxyRequestEvent;
import com.amazonaws.services.lambda.runtime.events.APIGatewayProxyResponseEvent;
import lombok.RequiredArgsConstructor;
import lombok.SneakyThrows;
import lombok.extern.slf4j.Slf4j;
import me.josephcosentino.meals.rest.mapper.Jackson;
import me.josephcosentino.meals.rest.model.ErrorResponse;
import software.amazon.awssdk.http.HttpStatusCode;
import software.amazon.awssdk.http.SdkHttpMethod;

import java.util.Map;

/**
 * Handler for requests to Lambda function.
 */
@Slf4j
@RequiredArgsConstructor
public class ResourceRequestHandler<T> implements RequestHandler<APIGatewayProxyRequestEvent, APIGatewayProxyResponseEvent> {

    private final ResourceHandler<T> resourceHandler;

    public APIGatewayProxyResponseEvent handleRequest(APIGatewayProxyRequestEvent input,
                                                      Context context) {
        try {
            return doHandleRequest(input, context);
        } catch (Exception e) {
            log.error("Handle failed", e);
            return internalServerError();
        }
    }

    private APIGatewayProxyResponseEvent doHandleRequest(APIGatewayProxyRequestEvent input,
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
    private APIGatewayProxyResponseEvent get(APIGatewayProxyRequestEvent input) {
        final var resource = resourceHandler.getById(getIdFromPath(input)).get();
        log.info("GET: {}", resource);
        if (resource == null) {
            return notFound();
        }
        return success(resource);
    }

    @SneakyThrows
    private APIGatewayProxyResponseEvent put(APIGatewayProxyRequestEvent input) {
        final var bodyResource = resourceHandler.asResource(input.getBody());
        var updatedResource = resourceHandler.update(bodyResource).get();
        log.info("PUT: {}", updatedResource);
        return success(updatedResource);
    }

    @SneakyThrows
    private APIGatewayProxyResponseEvent post(APIGatewayProxyRequestEvent input) {
        final var bodyResource = resourceHandler.asResource(input.getBody());
        var updatedResource = resourceHandler.create(bodyResource).get();
        log.info("POST: {}", updatedResource);
        return success(updatedResource);
    }

    @SneakyThrows
    private APIGatewayProxyResponseEvent delete(APIGatewayProxyRequestEvent input) {
        final var resource = resourceHandler.delete(getIdFromPath(input)).get();
        log.info("DELETE: {}", resource);
        if (resource == null) {
            return notFound();
        }
        return success(resource);
    }

    private String getIdFromPath(APIGatewayProxyRequestEvent input) {
        return input.getPathParameters().get("id");
    }


    // TODO put responses in rest-framework module?

    private APIGatewayProxyResponseEvent success(Object content) {
        final var resp = new APIGatewayProxyResponseEvent();
        resp.setBody(toJsonString(content));
        resp.setHeaders(Map.of("Content-Type", "application/json"));
        resp.setStatusCode(HttpStatusCode.OK);
        return resp;
    }

    private APIGatewayProxyResponseEvent notFound() {
        final var resp = new APIGatewayProxyResponseEvent();
        resp.setHeaders(Map.of("Content-Type", "application/json"));
        resp.setStatusCode(HttpStatusCode.NOT_FOUND);
        return resp;
    }

    private APIGatewayProxyResponseEvent badRequest() {
        final var resp = new APIGatewayProxyResponseEvent();
        resp.setBody(createErrorResponseBody("Bad Request"));
        resp.setHeaders(Map.of("Content-Type", "application/json"));
        resp.setStatusCode(HttpStatusCode.BAD_REQUEST);
        return resp;
    }

    private APIGatewayProxyResponseEvent internalServerError() {
        final var resp = new APIGatewayProxyResponseEvent();
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
