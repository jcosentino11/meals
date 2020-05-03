package me.josephcosentino.meals.tables;

import lombok.NonNull;
import lombok.RequiredArgsConstructor;
import me.josephcosentino.meals.client.DynamoDb;
import me.josephcosentino.meals.model.Recipe;
import me.josephcosentino.meals.util.Environment;
import software.amazon.awssdk.enhanced.dynamodb.DynamoDbAsyncTable;
import software.amazon.awssdk.enhanced.dynamodb.DynamoDbEnhancedAsyncClient;
import software.amazon.awssdk.enhanced.dynamodb.TableSchema;

import static software.amazon.awssdk.enhanced.dynamodb.mapper.StaticAttributeTags.primaryPartitionKey;

@RequiredArgsConstructor
public class RecipeTable {

    public static final String TABLE_NAME_ENV = "db_table_recipe_name";
    public static final String DEFAULT_TABLE_NAME = "recipes";

    public static final TableSchema<Recipe> SCHEMA =
            TableSchema.builder(Recipe.class)
                    .newItemSupplier(Recipe::new)
                    .addAttribute(String.class, a -> a.name("recipe_id")
                            .getter(Recipe::getId)
                            .setter(Recipe::setId)
                            .tags(primaryPartitionKey()))
                    .build();

    public static RecipeTable fromEnv() {
        final var tableName = Environment.get(TABLE_NAME_ENV, null);
        return new RecipeTable(
                DynamoDb.clientInstanceFromEnv(),
                tableName != null ? tableName : DEFAULT_TABLE_NAME
        );
    }

    private final DynamoDbAsyncTable<Recipe> table;

    public RecipeTable(@NonNull DynamoDbEnhancedAsyncClient client) {
        this(client, DEFAULT_TABLE_NAME);
    }

    public RecipeTable(@NonNull DynamoDbEnhancedAsyncClient client, @NonNull String tableName) {
        table = client.table(tableName, SCHEMA);
    }

    public DynamoDbAsyncTable<Recipe> get() {
        return table;
    }

}
