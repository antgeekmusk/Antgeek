package com.antgeek.common;

import com.antgeek.bean.Record;
import com.ververica.cdc.debezium.DebeziumDeserializationSchema;
import org.apache.flink.api.common.typeinfo.TypeInformation;
import org.apache.flink.util.Collector;
import org.apache.kafka.connect.source.SourceRecord;
import org.apache.kafka.connect.data.Struct;

public class RecordDeserializationSchema implements DebeziumDeserializationSchema<Record> {
    @Override
    public void deserialize(SourceRecord sourceRecord, Collector<Record> collector) throws Exception {
        //获取 database和 table
        String topic = sourceRecord.topic();
        //分隔符得写 \\. 不然就报错
        String[] strings = topic.split("\\.");
        String database = strings[1];
        String table = strings[2];

        //获取 data
        Struct value = (Struct) sourceRecord.value();
        Record record = new Record();

        // 只要after的数据(before是变更前的数据不需要)
        Struct after = value.getStruct("after");
        record.setId(after.getInt64("id"));
        record.setUserId(after.getString("user_id"));
        record.setRecord(after.getString("record"));

        // 输出数据
        collector.collect(record);
    }

    @Override
    public TypeInformation<Record> getProducedType() {
        return TypeInformation.of(Record.class);
    }
}
