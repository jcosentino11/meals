package me.josephcosentino.meals.modules;

import dagger.Component;
import me.josephcosentino.meals.handlers.RecipeHandler;

import javax.inject.Singleton;

@Singleton
@Component(modules = {
        DynamodbModule.class,
        RestModule.class
})
public interface RecipeHandlerComponent {

    RecipeHandler handler();

}
