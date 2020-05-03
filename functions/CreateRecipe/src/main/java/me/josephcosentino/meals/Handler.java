package me.josephcosentino.meals;

import com.amazonaws.services.lambda.runtime.Context;
import com.amazonaws.services.lambda.runtime.RequestHandler;
import lombok.extern.slf4j.Slf4j;
import me.josephcosentino.meals.tables.RecipeTable;

import java.net.Inet4Address;
import java.util.Map;

/**
 * Handler for requests to Lambda function.
 */
@Slf4j
public class Handler implements RequestHandler<Object, Object> {

    private static final CreateRecipe createRecipe;

    static {
        createRecipe = new CreateRecipe(RecipeTable.fromEnv().get());
    }

    public Object handleRequest(final Object input, final Context context) {
        // TODO handle actual input

        try {
            createRecipe.createRecipe();
            // TODO give actual response
            return new GatewayResponse("Success!", Map.of(), 200);
        } catch (Exception e) {
            log.error("Handle failed", e);
            return new GatewayResponse("Failure", Map.of(), 500);
        }
    }
}
