package com.antgeek.common;

import com.antgeek.bean.Record;
import org.apache.flink.configuration.Configuration;
import org.apache.flink.streaming.api.functions.sink.RichSinkFunction;

import java.sql.Connection;
import java.sql.DriverManager;
import java.sql.PreparedStatement;
import java.util.HashMap;
import java.util.Map;

public class HashSinkFunction extends RichSinkFunction<Record> {
    // TODO 这里可以优化一下,从链接池中获取conn,这里测试使用就使用了最简单的方法
    private Connection conn;// 数据库链接
    private Map<Integer,TableSink> sinkMap; // sink map
    private Integer splitNum=3;// 分表数量



    public HashSinkFunction(Integer splitNum) {
        this.splitNum = splitNum;
    }

    // open 方法仅在任务创建时执行一次
    @Override
    public void open(Configuration parameters) throws Exception {
        super.open(parameters);
        Class.forName("com.mysql.cj.jdbc.Driver");
        conn = DriverManager.getConnection("jdbc:mysql://127.0.0.1:3306/test2?useAffectedRows=true&useUnicode=true&characterEncoding=utf-8&useSSL=false&serverTimezone=Asia/Shanghai&allowMultiQueries=true","root","12345678");

        sinkMap=new HashMap<>();
        //  初始化sink map
        for (int i = 0; i < splitNum; i++) {

            String sql=" insert into record_"+i+"(id,user_id,record) values(?,?,?) " +
                    " ON DUPLICATE KEY UPDATE " +
                    " id=?,user_id=?,record=? ";

            sinkMap.put(i, new TableSink(conn.prepareStatement(sql)));
        }
    }

    // TODO 这里也可以优化一下,使用状态将数据攒批后输出,这也可以调高吞吐量
    //每来一数据就执行一次
    @Override
    public void invoke(Record value, Context context){
        try {
            //根据哈希值获取对应的sink对象
            TableSink tableSink = sinkMap.get(Math.abs(value.getUserId().hashCode())%splitNum);
            System.out.println(value);
            // 执行插入
            tableSink.write(value);
        }catch (Exception e){
            e.printStackTrace();
        }

    }

    //任务结束时关闭
    @Override
    public void close() throws Exception {
        super.close();
        if(conn!=null){
            conn.close();
        }
    }
    public static class TableSink{
        private PreparedStatement ps;

        public TableSink(PreparedStatement ps) {
            this.ps = ps;
        }
        public void write(Record r) throws Exception{
            ps.setLong(1,r.getId());
            ps.setString(2,r.getUserId());
            ps.setString(3,r.getRecord());

            ps.setLong(4,r.getId());
            ps.setString(5,r.getUserId());
            ps.setString(6,r.getRecord());

            ps.execute();
        }
    }
}
