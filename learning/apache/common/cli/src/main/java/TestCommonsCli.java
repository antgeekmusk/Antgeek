import org.apache.commons.cli.*;

public class TestCommonsCli {
    public static void main(String[] args) throws ParseException {
        // 定义阶段
        Options options = new Options();

        Option filePathOption = new Option("f", true, "文件路径");
        filePathOption.setRequired(false); // 文件路径为必传选项
        filePathOption.setArgName("string");
        Option daemonRunOption = new Option("d", false, "后台执行");
        Option helpOption = new Option("h", "help", false, "帮忙文档");

        options.addOption(filePathOption);
        options.addOption(daemonRunOption);
        options.addOption(helpOption);


        // 解析阶段
        CommandLineParser parser = new DefaultParser();
        CommandLine cmd = parser.parse(options, args);

        // 查询阶段
        if(args.length==0 || cmd.hasOption(helpOption)){
            HelpFormatter formatter = new HelpFormatter();
            formatter.printHelp(" ",options);
        }

        // 获取文件路径
        if(cmd.hasOption(filePathOption)){
            String filePath = cmd.getOptionValue(filePathOption);
            System.out.println("文件路径 : "+filePath);
        }

        // 后台运行
        if(cmd.hasOption(daemonRunOption)){
            System.out.println("后台运行");
        }

    }
}
