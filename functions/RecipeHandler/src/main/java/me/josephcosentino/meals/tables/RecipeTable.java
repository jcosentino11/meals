package me.josephcosentino.meals.tables;

import lombok.NonNull;
import lombok.RequiredArgsConstructor;
import lombok.SneakyThrows;
import me.josephcosentino.meals.client.DynamoDb;
import me.josephcosentino.meals.model.Recipe;
import me.josephcosentino.meals.util.Environment;
import software.amazon.awssdk.enhanced.dynamodb.DynamoDbAsyncTable;
import software.amazon.awssdk.enhanced.dynamodb.DynamoDbEnhancedAsyncClient;
import software.amazon.awssdk.enhanced.dynamodb.TableSchema;
import software.amazon.awssdk.enhanced.dynamodb.model.CreateTableEnhancedRequest;
import software.amazon.awssdk.services.dynamodb.model.ProvisionedThroughput;
import software.amazon.awssdk.services.dynamodb.model.ResourceInUseException;

import java.util.concurrent.ExecutionException;

import static software.amazon.awssdk.enhanced.dynamodb.mapper.StaticAttributeTags.primaryPartitionKey;
import static software.amazon.awssdk.enhanced.dynamodb.mapper.StaticAttributeTags.primarySortKey;

@RequiredArgsConstructor
public class RecipeTable {

    public static final String TABLE_NAME_ENV = "db_table_recipe_name";
    public static final String DEFAULT_TABLE_NAME = "recipes";

    public static final TableSchema<Recipe> SCHEMA = TableSchema.fromBean(Recipe.class);

    public static RecipeTable fromEnv() {
        final var tableName = Environment.get(TABLE_NAME_ENV, null);
        return new RecipeTable(
                DynamoDb.clientInstanceFromEnv(),
                tableName != null ? tableName : DEFAULT_TABLE_NAME,
                Boolean.parseBoolean(Environment.get("local", "false"))
        );
    }

    private final DynamoDbAsyncTable<Recipe> table;

    public RecipeTable(@NonNull DynamoDbEnhancedAsyncClient client) {
        this(client, DEFAULT_TABLE_NAME, false);
    }

    public RecipeTable(@NonNull DynamoDbEnhancedAsyncClient client, @NonNull String tableName, boolean autoCreate) {
        table = client.table(tableName, SCHEMA);
        if (autoCreate) {
            createTable();
        }
    }

    @SneakyThrows
    private void createTable() {
        try {
            table.createTable().get();
        } catch (ExecutionException e) {
            // ignore exceptions for already-created table
            if (e.getCause() == null || !(e.getCause() instanceof ResourceInUseException)) {
                throw e;
            }
        }
    }

    public DynamoDbAsyncTable<Recipe> get() {
        return table;
    }

}
