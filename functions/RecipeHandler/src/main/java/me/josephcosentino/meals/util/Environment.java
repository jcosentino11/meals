package me.josephcosentino.meals.util;

import lombok.AccessLevel;
import lombok.NoArgsConstructor;
import lombok.NonNull;

import java.util.Objects;

@NoArgsConstructor(access = AccessLevel.PRIVATE)
public final class Environment {

    public static String get(@NonNull String name) {
        var value = get(name, null);
        Objects.requireNonNull(value);
        return value;
    }

    public static String get(@NonNull String name, String defaultValue) {
        var value = System.getenv(name);
        return value != null && !value.isBlank() ? value : defaultValue;
    }

}
