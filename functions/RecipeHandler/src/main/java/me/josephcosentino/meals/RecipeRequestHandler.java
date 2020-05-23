package me.josephcosentino.meals;

import lombok.extern.slf4j.Slf4j;
import me.josephcosentino.meals.rest.handler.ResourceHandler;
import me.josephcosentino.meals.rest.handler.ResourceRequestHandler;

import javax.inject.Inject;

/**
 * Handler for requests to Lambda function.
 */
@Slf4j
public class RecipeRequestHandler extends ResourceRequestHandler<RecipeRequestHandler> {

    public RecipeRequestHandler(ResourceHandler<RecipeRequestHandler> resourceHandler) {
        super(resourceHandler);
    }

}
