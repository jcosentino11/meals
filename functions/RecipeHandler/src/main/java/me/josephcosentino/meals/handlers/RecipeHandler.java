package me.josephcosentino.meals.handlers;

import lombok.extern.slf4j.Slf4j;
import me.josephcosentino.meals.dynamodb.TableHandler;
import me.josephcosentino.meals.model.Recipe;
import software.amazon.awssdk.enhanced.dynamodb.DynamoDbAsyncTable;

import javax.inject.Inject;

// TODO make this generic, allow pluggable validation logic
@Slf4j
public class RecipeHandler extends TableHandler<Recipe> {

    @Inject
    public RecipeHandler(DynamoDbAsyncTable<Recipe> recipeTable) {
        super(recipeTable);
    }

    @Override
    public Class<Recipe> getSupportedResourceClass() {
        return Recipe.class;
    }

    @Override
    protected Recipe withId(Recipe recipe, String generatedId) {
        return recipe.toBuilder().id(generatedId).build();
    }
}
