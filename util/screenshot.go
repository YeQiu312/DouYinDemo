package util

import (
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func ExtractVideoCover(inputFile, outputFile string) error {
	// 使用 ffmpeg 获取视频时长
	cmdDuration := exec.Command("D:\\fttmp\\ffmpeg-2023-04-10-git-b18a9c2971-essentials_build\\ffmpeg-2023-04-10-git-b18a9c2971-essentials_build\\bin\\ffprobe", "-v", "error", "-show_entries", "format=duration", "-of", "default=noprint_wrappers=1:nokey=1", inputFile)
	durationBytes, err := cmdDuration.Output()
	if err != nil {
		log.Println("获取视频时长出错：", err)
		return err
	}

	durationStr := strings.TrimSpace(string(durationBytes))
	durationFloat, err := strconv.ParseFloat(durationStr, 64)
	if err != nil {
		log.Println("解析时长出错：", err)
		return err
	}

	// 使用 ffmpeg 智能选取一帧作为封面
	cmdCover := exec.Command("D:\\fttmp\\ffmpeg-2023-04-10-git-b18a9c2971-essentials_build\\ffmpeg-2023-04-10-git-b18a9c2971-essentials_build\\bin\\ffmpeg", "-y", "-ss", "00:00:01", "-i", inputFile, "-t", strconv.FormatFloat(durationFloat, 'f', -1, 64), "-vf", "select='eq(pict_type,I)'", "-frames:v", "1", "-f", "image2", outputFile)
	cmdCover.Stderr = os.Stderr // 设置标准错误输出
	if err := cmdCover.Run(); err != nil {
		log.Println("截取封面出错：", err.Error())
		return err
	}
	return nil
}
