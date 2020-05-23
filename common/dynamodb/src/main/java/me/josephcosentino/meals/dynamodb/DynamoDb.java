package me.josephcosentino.meals.dynamodb;

import lombok.extern.slf4j.Slf4j;
import software.amazon.awssdk.enhanced.dynamodb.DynamoDbEnhancedAsyncClient;
import software.amazon.awssdk.regions.Region;
import software.amazon.awssdk.services.dynamodb.DynamoDbAsyncClient;

import java.net.URI;


// TODO refactor to module
@Slf4j
public class DynamoDb {

    public static DynamoDbBuilder builder() {
        return new DynamoDbBuilder();
    }

    public static class DynamoDbBuilder {

        private String localEndpoint;
        private String region;

        public DynamoDbBuilder localEndpoint(String localEndpoint) {
            this.localEndpoint = localEndpoint;
            return this;
        }

        public DynamoDbBuilder region(String region) {
            this.region = region;
            return this;
        }

        public DynamoDbEnhancedAsyncClient build() {
            final var stdBuilder = DynamoDbAsyncClient.builder();

            if (localEndpoint != null && !localEndpoint.isBlank()) {
                log.info("using local endpoint: {}", localEndpoint);
                stdBuilder.endpointOverride(URI.create(localEndpoint));
            }

            if (region != null && !region.isBlank()) {
                log.info("using region: {}", region);
                stdBuilder.region(Region.of(region));
            }

            final var stdClient = stdBuilder.build();

            return DynamoDbEnhancedAsyncClient.builder()
                    .dynamoDbClient(stdClient)
                    .build();
        }
    }
}
