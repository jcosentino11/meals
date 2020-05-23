package me.josephcosentino.meals.modules;

import dagger.Module;
import dagger.Provides;
import me.josephcosentino.meals.dynamodb.DynamoDb;
import me.josephcosentino.meals.model.Recipe;
import me.josephcosentino.meals.dynamodb.Table;
import me.josephcosentino.meals.rest.util.Environment;
import software.amazon.awssdk.enhanced.dynamodb.DynamoDbAsyncTable;
import software.amazon.awssdk.enhanced.dynamodb.DynamoDbEnhancedAsyncClient;

import javax.inject.Singleton;

@Module
public class DynamodbModule {

    public static final String PROP_DB_LOCAL_ENDPOINT = "db_localEndpoint";
    public static final String PROP_DB_REGION = "db_region";
    public static final String PROP_RECIPE_TABLE_NAME = "db_table_recipe_name";
    public static final String PROP_RECIPE_TABLE_AUTOCREATE = "local";

    @Provides
    @Singleton
    DynamoDbAsyncTable<Recipe> recipeDynamoDbTable(Table<Recipe> recipeTable) {
        return recipeTable.get();
    }

    @Provides
    @Singleton
    Table<Recipe> recipeTable(DynamoDbEnhancedAsyncClient dynamoDbClient) {
        return new Table<>(
                dynamoDbClient,
                Recipe.class,
                Environment.get(PROP_RECIPE_TABLE_NAME, "recipes"),
                Environment.get(PROP_RECIPE_TABLE_AUTOCREATE, false)
        );
    }

    @Provides
    @Singleton
    DynamoDbEnhancedAsyncClient dynamoDbClient() {
        return DynamoDb.builder()
                .localEndpoint(Environment.get(PROP_DB_LOCAL_ENDPOINT, null))
                .region(Environment.get(PROP_DB_REGION, null))
                .build();
    }

}
