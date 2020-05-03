package me.josephcosentino.meals;

import lombok.RequiredArgsConstructor;
import me.josephcosentino.meals.model.Recipe;
import software.amazon.awssdk.enhanced.dynamodb.DynamoDbAsyncTable;

@RequiredArgsConstructor
public class CreateRecipe {

    private int count;

    private final DynamoDbAsyncTable<Recipe> recipeTable;

    public void createRecipe() throws Exception { // TODO input
        recipeTable
                .putItem(Recipe.builder()
                        .id(String.valueOf(count++))
                        .build())
                .get();
    }

}
