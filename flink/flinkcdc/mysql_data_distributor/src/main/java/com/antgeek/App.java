package com.antgeek;

import com.antgeek.bean.Record;
import com.antgeek.common.HashSinkFunction;
import com.antgeek.common.RecordDeserializationSchema;
import com.ververica.cdc.connectors.mysql.MySqlSource;
import com.ververica.cdc.debezium.StringDebeziumDeserializationSchema;
import org.apache.flink.api.common.functions.MapFunction;
import org.apache.flink.api.common.typeinfo.BasicTypeInfo;
import org.apache.flink.api.common.typeinfo.TypeInformation;
import org.apache.flink.configuration.Configuration;
import org.apache.flink.configuration.RestOptions;
import org.apache.flink.connector.jdbc.JdbcConnectionOptions;
import org.apache.flink.connector.jdbc.JdbcExecutionOptions;
import org.apache.flink.connector.jdbc.JdbcSink;
import org.apache.flink.streaming.api.CheckpointingMode;
import org.apache.flink.streaming.api.datastream.DataStreamSink;
import org.apache.flink.streaming.api.environment.CheckpointConfig;
import org.apache.flink.streaming.api.functions.sink.SinkFunction;
import org.apache.flink.streaming.api.functions.source.SourceFunction;
import org.apache.flink.streaming.api.scala.DataStream;
import org.apache.flink.streaming.api.scala.StreamExecutionEnvironment;
import org.apache.flink.table.api.EnvironmentSettings;
import org.apache.flink.table.api.bridge.java.StreamTableEnvironment;

public class App {
    public static void main(String[] args) {
        Configuration conf = new Configuration();
        conf.setString(RestOptions.BIND_PORT,"8081");
        // 获取执行环境
        StreamExecutionEnvironment env = StreamExecutionEnvironment.createLocalEnvironmentWithWebUI(conf);
        env.setParallelism(4);

        // 配置checkpoint
        CheckpointConfig ckConfig = env.getCheckpointConfig();
        ckConfig.setCheckpointInterval(4000L);
        ckConfig.setCheckpointingMode(CheckpointingMode.EXACTLY_ONCE);
        ckConfig.setMaxConcurrentCheckpoints(1);
        ckConfig.setMinPauseBetweenCheckpoints(500);



        // 获取数据源
        SourceFunction<Record> source = MySqlSource.<Record>builder()
                .hostname("127.0.0.1")
                .port(3306)
                .username("root")
                .password("12345678")
                .databaseList("test")
                .tableList("test.record")
                .deserializer(new RecordDeserializationSchema())
                .build();

        //打印数据
        DataStream<Record> ds = env.addSource(source, TypeInformation.of(Record.class));

        // 根据hash sink到不同的表
        ds.addSink(new HashSinkFunction(3));

        // 执行
        env.execute();


    }
}
