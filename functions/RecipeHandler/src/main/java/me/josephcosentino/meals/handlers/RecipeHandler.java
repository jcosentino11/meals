package me.josephcosentino.meals.handlers;

import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import me.josephcosentino.meals.model.Recipe;
import software.amazon.awssdk.enhanced.dynamodb.DynamoDbAsyncTable;

import java.util.concurrent.CompletableFuture;

// TODO make this generic, allow pluggable validation logic
@Slf4j
@RequiredArgsConstructor
public class RecipeHandler implements ResourceHandler<Recipe> {

    private final DynamoDbAsyncTable<Recipe> recipeTable;

    @Override
    public CompletableFuture<Recipe> getById(String id) {
        log.info("getById - id={}", id);
        return recipeTable.getItem(keyFromId(id));
    }

    @Override
    public CompletableFuture<Recipe> create(Recipe recipe) {
        log.info("create - recipe={}", recipe);
        // TODO validation
        final var recipeWithId = recipe.toBuilder().id(newRandomId()).build();
        return recipeTable.putItem(recipeWithId).thenApply(none -> recipeWithId);
    }

    @Override
    public CompletableFuture<Recipe> update(Recipe recipe) {
        log.info("update - recipe={}", recipe);
        // TODO validation
        return recipeTable.updateItem(recipe);
    }

    @Override
    public CompletableFuture<Recipe> delete(String id) {
        log.info("delete - id={}", id);
        return recipeTable.deleteItem(keyFromId(id));
    }

    @Override
    public Class<Recipe> getSupportedResourceClass() {
        return Recipe.class;
    }

}
