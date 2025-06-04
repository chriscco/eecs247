import java.util.*;
import java.util.concurrent.TimeoutException;

import org.apache.commons.collections.map.UnmodifiableMap;
import org.apache.spark.api.java.function.FlatMapFunction;
import org.apache.spark.sql.Dataset;
import org.apache.spark.sql.Encoders;
import org.apache.spark.sql.Row;
import org.apache.spark.sql.SparkSession;
import org.apache.spark.sql.streaming.StreamingQuery;
import org.apache.spark.sql.streaming.StreamingQueryException;

public class WordCounter {
    private SparkSession sparkSession;
    private Dataset<Row> kafkaDF;
    private Map<String, Long> map;

    /**
     * Constructor
     */
    public WordCounter() {
        sparkSession = SparkSession
                .builder()
                .appName("SparkWordCounter")
                .master("local[*]")
                .getOrCreate();
        map = new HashMap<>();
    }

    /**
     * create data frame
     */
    private void createDataframe() {
        kafkaDF = sparkSession
                .readStream()
                .format("kafka")
                .option("kafka.bootstrap.servers", "localhost:9092")
                .option("subscribe", "kafka_one")
                .load()
                .selectExpr("value AS STRING");
    }

    /**
     * do count
     */
    public void counter() {
        try {
            createDataframe();

            long start = System.currentTimeMillis();
            Dataset<String> words = kafkaDF
                    .as(Encoders.STRING())
                    .flatMap((FlatMapFunction<String, String>) x -> {
                        if ("__END__".equals(x)) {
                            long end = System.currentTimeMillis();
                            System.out.println("Time Taken (Spark Streaming): " + (end - start) + "ms");
                        }
                        return Arrays.asList(x.split(" ")).iterator();
                    }, Encoders.STRING());
            Dataset<Row> wordCount = words.groupBy("value").count();
            StreamingQuery query = wordCount
                    .writeStream()
                    .outputMode("complete")
                    .foreachBatch((batchDf, batchId) -> {
                        List<Row> list = batchDf.collectAsList();

                        for (Row r : list) {
                            String word = r.getString(0);
                            Long count = r.getLong(1);
                            map.put(word, count);
                        }

                    })
                    .start();
            query.awaitTermination();
        } catch (TimeoutException | StreamingQueryException e) {
            System.out.println(e);
        }
    }

    /**
     * get count result
     * @return map
     */
    public Map<String, Long> getMap() {
        return Collections.unmodifiableMap(map);
    }
}
