package server;

import WordCountService.ServerProto;
import WordCountService.WordCountGrpc;
import io.grpc.stub.StreamObserver;

public class WordCountServiceImpl extends WordCountGrpc.WordCountImplBase {
    @Override
    public void wordCount(ServerProto.WordCountRequest request, StreamObserver<ServerProto.WordCountResponse> responseObserver) {
        ServerProto.WordCountResponse response =
    }
}