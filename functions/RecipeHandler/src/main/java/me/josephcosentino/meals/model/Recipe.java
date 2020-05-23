package me.josephcosentino.meals.model;

import lombok.*;
import software.amazon.awssdk.enhanced.dynamodb.mapper.annotations.DynamoDbBean;
import software.amazon.awssdk.enhanced.dynamodb.mapper.annotations.DynamoDbPartitionKey;

@Data
@Builder(toBuilder = true)
@AllArgsConstructor
@NoArgsConstructor
@DynamoDbBean
public class Recipe {

    @Getter(onMethod_ = {@DynamoDbPartitionKey})
    private String id;
    private String name;

}
