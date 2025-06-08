package server;

import java.security.MessageDigest;
import java.security.NoSuchAlgorithmException;
import java.util.Map;
import java.util.concurrent.ConcurrentHashMap;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.concurrent.atomic.AtomicInteger;

import WordCountService.ServerProto;
import WordCountService.WordCountGrpc;
import io.grpc.stub.StreamObserver;

public class WordCountServiceImpl extends WordCountGrpc.WordCountImplBase {
    private final ConcurrentHashMap<String, AtomicInteger> result = new ConcurrentHashMap<>();
    private final ExecutorService executor = Executors.newFixedThreadPool(4);

    /**
     * count words in multi-thread mode
     * @param responseObserver responseObserver
     * @return StreamObserver
     */
    @Override
    public StreamObserver<ServerProto.WordCountRequest> wordCount(StreamObserver<ServerProto.WordCountResponse> responseObserver) {
        return new StreamObserver<ServerProto.WordCountRequest>() {
            @Override
            public void onNext(ServerProto.WordCountRequest wordCountRequest) {
                executor.submit(() -> {
                    String input = wordCountRequest.getRequestMessage();
                    String key = keyGen(input);
                    String[] words = input.split("\\s+");
                    System.out.println("getRequestMessage: " + wordCountRequest.getRequestMessage() +
                            "Key: " + key);

                    for (String word : words) {
                        result.computeIfAbsent(word, k -> new AtomicInteger(0))
                                .incrementAndGet();

                        ServerProto.WordCountResponse.Builder responseBuilder = ServerProto.WordCountResponse.newBuilder();
                        for (Map.Entry<String, AtomicInteger> entry : result.entrySet()) {
                            ServerProto.WordCountResponse.WordCountResult temp =
                                    ServerProto.WordCountResponse.WordCountResult
                                            .newBuilder()
                                            .setWord(entry.getKey())
                                            .setCount(entry.getValue().get())
                                            .build();
                            responseBuilder.addCt(temp);
                            responseBuilder.setKey(key);
                        }
                        // System.out.println("sending one response");
                        responseObserver.onNext(responseBuilder.build());
                    }
                });
            }

            @Override
            public void onError(Throwable throwable) {
                System.out.println("Error in gRPC Streaming: " + throwable.getMessage());
            }

            @Override
            public void onCompleted() {
                responseObserver.onCompleted();
            }
        };
    }

    private String keyGen(String str) {
        try {
            MessageDigest sha1 = MessageDigest.getInstance("SHA-1");

            byte[] hash = sha1.digest(str.getBytes());
            String hashPrefix = bytesToHex(hash).substring(0, 16);
            return String.format("word_count:%s", hashPrefix);
        } catch (NoSuchAlgorithmException e) {
            return "";
        }
    }

    public static String bytesToHex(byte[] bytes) {
        StringBuilder hexString = new StringBuilder();
        for (byte b : bytes) {
            hexString.append(String.format("%02x", b));
        }
        return hexString.toString();
    }
}