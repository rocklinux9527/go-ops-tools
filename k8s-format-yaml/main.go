package main

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"k8s-format-yaml/logger"
	"log"
	"os"
)

type InputArgs struct {
	SourcePath string `flag:"sourcePath, SourcePath"`
	TargetPath string `flag:"targetPath,TargetPath"`
}

func k8sYamlAnalyzeFormat(sourcePath string, destPath string) {
	log.Printf("转化函数入参数: 文件输入路径: %s   文件输出路径:%s", sourcePath, destPath)

	if sourcePath == "" && destPath == "" {
		fmt.Println("错误：没有配置输出yaml来源路径和输出路径,程序终止")
		os.Exit(1)
	}
	// 读取原始 yaml 文件
	data, err := ioutil.ReadFile(sourcePath)
	if err != nil {
		log.Fatalf("读取文件失败：%v", err)
	}

	// 检查文件是否为空
	if len(data) == 0 {
		log.Printf("警告：文件名称为空：%s", sourcePath)
		return
	}

	// 将 yaml 解析为 map[string]interface{} 类型
	yamlData := make(map[string]interface{})
	err = yaml.Unmarshal(data, &yamlData)
	if err != nil {
		log.Fatalf("解析 yaml 失败：%v", err)
	}

	delete(yamlData, "status") //	删除 status 节点下的所有内容
	log.Printf("删除k8s status 字段中所有字段 sucess")

	nsMetadata, ok := yamlData["metadata"].(map[interface{}]interface{}) // 清理crd添加的属性

	if ok {
		// 列出要删除的键 删除多个同级别的删除的键
		keysToDelete := []string{"managedFields", "resourceVersion", "selfLink", "uid", "nsMetadata", "creationTimestamp", "generation", "annotations"}

		// 遍历要删除的键，并逐一删除
		for _, key := range keysToDelete {
			delete(nsMetadata, key)
			log.Printf("删除k8s metadata 字段中的 %s 字段成功", key)
		}
	}


	spec, specOk := yamlData["spec"].(map[interface{}]interface{})
	if specOk {

		template, templateOk := spec["template"].(map[interface{}]interface{})
		if templateOk {
			metadata, metadataOk := template["metadata"].(map[interface{}]interface{})
			if metadataOk {
				delete(metadata, "creationTimestamp") // 清理spec.template.metadata.creationTimestamp 的属性
				log.Printf("删除k8s spec.template.metadata 字段中的 creationTimestamp sucess")
			}
		}
	}

	specV2, specOkV2 := yamlData["spec"].(map[interface{}]interface{})
	if specOkV2 {
		delete(specV2, "clusterIP")
		log.Printf("删除k8s spec.clusterIP 字段中的 ip地址")
	}

	// 将 map 转换回 yaml 格式
	newData, err := yaml.Marshal(&yamlData)
	if err != nil {
		log.Fatalf("生成 yaml 失败：%v", err)
	}

	// 输出新的 yaml 文件
	err = ioutil.WriteFile(destPath, newData, 0644)
	if err != nil {
		log.Fatalf("写入文件失败：%v", err)
	}

	fmt.Println("新的 yaml 文件已生成！")
}

func main() {
	logRotation := logger.LogBck()
	log.SetOutput(logRotation)
	defer logRotation.Close() // 延迟关闭logRotation
	log.Printf("start 开始转化k8s yaml")
	inputArgs := InputArgs{}
	// 解析命令行参数，并将其分配给结构体字段
	flag.StringVar(&inputArgs.SourcePath, "sourcePath", "", "Source Path")
	flag.StringVar(&inputArgs.TargetPath, "targetPath", "", "Target Path")

	// 自定义帮助信息
	flag.Usage = func() {
		fmt.Println("Usage: program [options]")
		fmt.Println("Options:")
		flag.PrintDefaults()
		fmt.Println("For more information")
	}
	// 解析命令行参数
	flag.Parse()

	switch {
	case inputArgs.SourcePath == "":
		fmt.Println("错误：必须提供k8s yaml 的完整来源路径(包括文件名称),文件名称")
		log.Printf("没有参数 来源路径参数缺少 SourcePath ")
		flag.PrintDefaults()
		return

	case inputArgs.TargetPath == "":
		fmt.Println("Error：必须提供k8s yaml 转化以后完整输出路径(包括文件名称) 文件名称 ")
		log.Printf("没有参数 输出路径参数缺少 TargetPath ")
		flag.PrintDefaults()
		return

	default:
		// 打印结构体中的字段值
		fmt.Println("SourcePath:", inputArgs.SourcePath)
		fmt.Println("TargetPath:", inputArgs.TargetPath)
		log.Printf("用户输入参数: 文件输入路径: %s   文件输出路径:%s", inputArgs.SourcePath, inputArgs.TargetPath)
		k8sYamlAnalyzeFormat(inputArgs.SourcePath, inputArgs.TargetPath)
		log.Printf("end 结束转化k8s yaml")
	}
}
