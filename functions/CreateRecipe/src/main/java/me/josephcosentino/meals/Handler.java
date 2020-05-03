package me.josephcosentino.meals;

import com.amazonaws.services.lambda.runtime.Context;
import com.amazonaws.services.lambda.runtime.RequestHandler;
import com.amazonaws.services.lambda.runtime.events.APIGatewayV2ProxyRequestEvent;
import com.amazonaws.services.lambda.runtime.events.APIGatewayV2ProxyResponseEvent;
import lombok.extern.slf4j.Slf4j;
import me.josephcosentino.meals.tables.RecipeTable;
import software.amazon.awssdk.http.HttpStatusCode;

/**
 * Handler for requests to Lambda function.
 */
@Slf4j
public class Handler implements RequestHandler<APIGatewayV2ProxyRequestEvent, APIGatewayV2ProxyResponseEvent> {

    private static final CreateRecipe createRecipe;

    static {
        createRecipe = new CreateRecipe(RecipeTable.fromEnv().get());
    }

    public APIGatewayV2ProxyResponseEvent handleRequest(APIGatewayV2ProxyRequestEvent input,
                                                        Context context) {
        // TODO handle actual input
        try {
            createRecipe.createRecipe();
            // TODO give actual response
            return success();
        } catch (Exception e) {
            log.error("Handle failed", e);
            return failure();
        }
    }

    private APIGatewayV2ProxyResponseEvent success() {
        final var resp = new APIGatewayV2ProxyResponseEvent();
        resp.setBody("Success!");
        resp.setStatusCode(HttpStatusCode.OK);
        return resp;
    }

    private APIGatewayV2ProxyResponseEvent failure() {
        final var resp = new APIGatewayV2ProxyResponseEvent();
        resp.setBody("Failure!");
        resp.setStatusCode(HttpStatusCode.INTERNAL_SERVER_ERROR);
        return resp;
    }
}
