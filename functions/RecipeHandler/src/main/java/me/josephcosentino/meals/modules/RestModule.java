package me.josephcosentino.meals.modules;

import dagger.Module;
import dagger.Provides;
import me.josephcosentino.meals.handlers.RecipeHandler;
import me.josephcosentino.meals.model.Recipe;
import software.amazon.awssdk.enhanced.dynamodb.DynamoDbAsyncTable;

import javax.inject.Singleton;

@Module(includes = DynamodbModule.class)
public class RestModule {

    @Provides
    @Singleton
    RecipeHandler recipeDynamoDbTable(DynamoDbAsyncTable<Recipe> recipeDynamoDbTable) {
        return new RecipeHandler(recipeDynamoDbTable);
    }

}
