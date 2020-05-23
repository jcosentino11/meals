package me.josephcosentino.meals.dynamodb;

import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import me.josephcosentino.meals.rest.handler.ResourceHandler;
import software.amazon.awssdk.enhanced.dynamodb.DynamoDbAsyncTable;
import software.amazon.awssdk.enhanced.dynamodb.Key;

import java.util.concurrent.CompletableFuture;

// TODO  allow pluggable validation logic
@Slf4j
@RequiredArgsConstructor
public abstract class TableHandler<T> implements ResourceHandler<T> {

    private final DynamoDbAsyncTable<T> table;

    @Override
    public CompletableFuture<T> getById(String id) {
        log.info("getById - id={}", id);
        return table.getItem(keyFromId(id));
    }

    @Override
    public CompletableFuture<T> create(T resource) {
        log.info("create - resource={}", resource);
        // TODO validation
        final var recipeWithId = withId(resource, newRandomId());
        return table.putItem(recipeWithId).thenApply(none -> recipeWithId);
    }

    @Override
    public CompletableFuture<T> update(T resource) {
        log.info("update - resource={}", resource);
        // TODO validation
        return table.updateItem(resource);
    }

    @Override
    public CompletableFuture<T> delete(String id) {
        log.info("delete - id={}", id);
        return table.deleteItem(keyFromId(id));
    }

    abstract protected T withId(T resource, String generatedId);

    private Key keyFromId(String id) {
        return Key.builder().partitionValue(id).build();
    }
}
