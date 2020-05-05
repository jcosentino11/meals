package me.josephcosentino.meals.handlers;

import lombok.AccessLevel;
import lombok.RequiredArgsConstructor;
import me.josephcosentino.meals.handlers.RecipeHandler;
import me.josephcosentino.meals.handlers.ResourceHandler;
import me.josephcosentino.meals.model.Recipe;
import me.josephcosentino.meals.tables.RecipeTable;

import java.util.Map;
import java.util.function.Supplier;

@RequiredArgsConstructor(access = AccessLevel.PRIVATE)
public final class ResourceHandlers {

    // TODO find a way to register automatically?
    private static final Map<String, Supplier<ResourceHandler<?>>> RESOURCE_HANDLERS = Map.of(
            ResourceHandler.getSupportedResourceName(Recipe.class), () -> new RecipeHandler(RecipeTable.fromEnv().get())
    );

    public static Map<String, Supplier<ResourceHandler<?>>> all() {
        return RESOURCE_HANDLERS;
    }

    public static Supplier<ResourceHandler<?>> get(String resourceName) {
        return all().get(resourceName);
    }

}
