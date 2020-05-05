package me.josephcosentino.meals.client;

import lombok.extern.slf4j.Slf4j;
import me.josephcosentino.meals.util.Environment;
import software.amazon.awssdk.enhanced.dynamodb.DynamoDbEnhancedAsyncClient;
import software.amazon.awssdk.regions.Region;
import software.amazon.awssdk.services.dynamodb.DynamoDbAsyncClient;

import java.net.URI;


@Slf4j
public class DynamoDb {

    public static final String DB_LOCAL_ENDPOINT_ENV = "db_localEndpoint";
    public static final String DB_REGION_ENV = "db_region";

    private static DynamoDbEnhancedAsyncClient CLIENT_INSTANCE;

    public static DynamoDbEnhancedAsyncClient clientInstanceFromEnv() {
        if (CLIENT_INSTANCE == null) {
            CLIENT_INSTANCE = newClientFromEnv();
        }
        return CLIENT_INSTANCE;
    }

    public static DynamoDbEnhancedAsyncClient newClientFromEnv() {
        final var builder = builder();

        log.info("Creating new dynamodb client from env");

        final var localEndpoint = Environment.get(DB_LOCAL_ENDPOINT_ENV, null);
        if (localEndpoint != null) {
            log.info("using local endpoint: {}", localEndpoint);
            builder.localEndpoint(localEndpoint);
        }

        final var region = Environment.get(DB_REGION_ENV, null);
        if (region != null) {
            log.info("using region: {}", region);
            builder.region(region);
        }

        return builder.build();
    }

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
                stdBuilder.endpointOverride(URI.create(localEndpoint));
            }

            if (region != null && !region.isBlank()) {
                stdBuilder.region(Region.of(region));
            }

            final var stdClient = stdBuilder.build();

            return DynamoDbEnhancedAsyncClient.builder()
                    .dynamoDbClient(stdClient)
                    .build();
        }
    }
}
