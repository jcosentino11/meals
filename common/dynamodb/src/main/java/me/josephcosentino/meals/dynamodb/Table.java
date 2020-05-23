package me.josephcosentino.meals.dynamodb;

import lombok.NonNull;
import lombok.SneakyThrows;
import software.amazon.awssdk.enhanced.dynamodb.DynamoDbAsyncTable;
import software.amazon.awssdk.enhanced.dynamodb.DynamoDbEnhancedAsyncClient;
import software.amazon.awssdk.enhanced.dynamodb.TableSchema;
import software.amazon.awssdk.services.dynamodb.model.ResourceInUseException;

import java.util.concurrent.ExecutionException;

public class Table<T> {

    private final DynamoDbAsyncTable<T> table;

    public Table(@NonNull DynamoDbEnhancedAsyncClient client,
                 Class<T> schema,
                 @NonNull String tableName,
                 boolean autoCreate) {
        table = client.table(tableName, TableSchema.fromBean(schema));
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

    public DynamoDbAsyncTable<T> get() {
        return table;
    }

}
