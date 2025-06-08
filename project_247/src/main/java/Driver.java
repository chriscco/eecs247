import io.grpc.Server;
import io.grpc.ServerBuilder;
import server.WordCountServiceImpl;

import java.io.IOException;

public class Driver {
    public static void main(String[] args) throws IOException, InterruptedException {
        WordCountServiceImpl wordCountService = new WordCountServiceImpl();
        Server server = ServerBuilder
                .forPort(9091)
                .addService(wordCountService)
                .maxInboundMessageSize(1024*1024*20)
                .build()
                .start();
        System.out.println("running on 9091");
        Runtime.getRuntime().addShutdownHook(new Thread(() -> {
            System.out.println("shutting down grpc server");
            server.shutdown();
        }));
        server.awaitTermination();
    }
}
