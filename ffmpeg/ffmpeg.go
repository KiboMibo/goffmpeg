package ffmpeg

import (
  "github.com/KiboMibo/goffmpeg/utils"
  "os/exec"
  "bytes"
  "strings"
)

type Configuration struct {
  FfmpegBin  string
  FfprobeBin string
  ExecCmd    string
  ExecArgs   string
}

var FFConfig *Configuration

func Configure() (*Configuration, error) {

  if FFConfig == nil {

    FFConfig = new(Configuration)

    var outFFmpeg bytes.Buffer
    var outProbe bytes.Buffer

    execCmd := utils.GetExec()
    execFFmpegCommand := utils.GetFFmpegExec()
    execFFprobeCommand := utils.GetFFprobeExec()
    execArgs := utils.GetExecArgs()

    cmdFFmpeg := exec.Command(execCmd, execArgs, execFFmpegCommand)
    cmdProbe := exec.Command(execCmd, execArgs, execFFprobeCommand)

    cmdFFmpeg.Stdout = &outFFmpeg
    cmdProbe.Stdout = &outProbe

    err := cmdFFmpeg.Start()
    if err != nil {
      return nil, err
    }

    _, err = cmdFFmpeg.Process.Wait()
    if err != nil {
      return nil, err
    }

    err = cmdProbe.Start()
    if err != nil {
      return nil, err
    }

    _, err = cmdProbe.Process.Wait()
    if err != nil {
      return nil, err
    }

    ffmpeg := strings.Replace(outFFmpeg.String(), "\n", "", -1)
    fprobe := strings.Replace(outProbe.String(), "\n", "", -1)

    FFConfig.FfmpegBin = ffmpeg
    FFConfig.FfprobeBin = fprobe
    FFConfig.ExecCmd = execCmd
    FFConfig.ExecArgs = execArgs

  }

  return FFConfig, nil

}
