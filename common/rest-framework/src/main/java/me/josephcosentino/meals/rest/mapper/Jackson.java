package me.josephcosentino.meals.rest.mapper;

import com.fasterxml.jackson.databind.ObjectMapper;

public class Jackson {

    private static final ObjectMapper INSTANCE;

    static {
        INSTANCE = new ObjectMapper();
    }

    public static ObjectMapper getObjectMapper() {
        return INSTANCE;
    }

}
