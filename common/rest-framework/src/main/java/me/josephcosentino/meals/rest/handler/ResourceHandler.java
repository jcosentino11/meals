package me.josephcosentino.meals.rest.handler;

import lombok.SneakyThrows;
import me.josephcosentino.meals.rest.mapper.Jackson;

import java.util.UUID;
import java.util.concurrent.CompletableFuture;

public interface ResourceHandler<T> {

    CompletableFuture<T> getById(String id);

    CompletableFuture<T> create(T recipe);

    CompletableFuture<T> update(T recipe);

    CompletableFuture<T> delete(String id);

    Class<T> getSupportedResourceClass();

    default String getSupportedResourceName() {
        return getSupportedResourceClass().getSimpleName().toLowerCase();
    }

    static <T> String getSupportedResourceName(Class<T> clazz) {
        return clazz.getSimpleName().toLowerCase();
    }

    @SneakyThrows
    default T asResource(String jsonValue) {
        return Jackson.getObjectMapper().readValue(jsonValue, getSupportedResourceClass());
    }

    default String newRandomId() {
        return UUID.randomUUID().toString();
    }

}
