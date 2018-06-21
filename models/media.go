package models

import (
  "strings"
  "strconv"
  "fmt"
  "reflect"
)

type Mediafile struct {
    aspect                string
    resolution            string
    videoBitRate          int
    videoBitRateTolerance int
    videoMaxBitRate       int
    videoMinBitrate       int
    videoCodec            string
    frameRate             int
    maxKeyframe           int
    minKeyframe           int
    keyframeInterval      int
    audioCodec            string
    audioBitrate          int
    audioChannels         int
    bufferSize            int
    threads               int
    preset                string
    tune                  string
    audioProfile          string
    videoProfile          string
    target                string
    duration              string
    durationInput         string
    seekTime              string
    quality               int
    muxDelay              string
    seekUsingTsInput      bool
    seekTimeInput         string
    inputPath             string
    outputPath            string
    outputFormat          string
    copyTs                bool
    nativeFramerateInput  bool
    rtmpLive              string
    hlsPlaylistType       string
    hlsSegmentDuration    int
    metadata              Metadata
    filter                string
}

/*** SETTERS ***/
func (m *Mediafile) SetFilter(v string){
  m.filter = v
}

func (m *Mediafile) SetAspect(v string) {
  m.aspect = v
}

func (m *Mediafile) SetResolution(v string) {
  m.resolution = v
}

func (m *Mediafile) SetVideoBitRate(v int) {
  m.videoBitRate = v
}

func (m *Mediafile) SetVideoBitRateTolerance(v int) {
  m.videoBitRateTolerance = v
}

func (m *Mediafile) SetVideoMaxBitrate(v int) {
  m.videoMaxBitRate = v
}

func (m *Mediafile) SetVideoMinBitRate(v int) {
  m.videoMinBitrate = v
}

func (m *Mediafile) SetVideoCodec(v string) {
  m.videoCodec = v
}

func (m *Mediafile) SetFrameRate(v int) {
  m.frameRate = v
}

func (m *Mediafile) SetMaxKeyFrame(v int) {
  m.maxKeyframe = v
}

func (m *Mediafile) SetMinKeyFrame(v int) {
  m.minKeyframe = v
}

func (m *Mediafile) SetKeyframeInterval(v int) {
  m.keyframeInterval = v
}

func (m *Mediafile) SetAudioCodec(v string) {
  m.audioCodec = v
}

func (m *Mediafile) SetAudioBitRate(v int) {
  m.audioBitrate = v
}

func (m *Mediafile) SetAudioChannels(v int) {
  m.audioChannels = v
}

func (m *Mediafile) SetBufferSize(v int) {
  m.bufferSize = v
}

func (m *Mediafile) SetThreads(v int) {
  m.threads = v
}

func (m *Mediafile) SetPreset(v string) {
  m.preset = v
}

func (m *Mediafile) SetTune(v string) {
  m.tune = v
}

func (m *Mediafile) SetAudioProfile(v string) {
  m.audioProfile = v
}

func (m *Mediafile) SetVideoProfile(v string) {
  m.videoProfile = v
}

func (m *Mediafile) SetDuration(v string) {
  m.duration = v
}

func (m *Mediafile) SetDurationInput(v string) {
  m.durationInput = v
}

func (m *Mediafile) SetSeekTime(v string) {
  m.seekTime = v
}

func (m *Mediafile) SetSeekTimeInput(v string) {
  m.seekTimeInput = v
}

func (m *Mediafile) SetQuality(v int) {
  m.quality = v
}

func (m *Mediafile) SetSeekUsingTsInput(val bool) {
  m.seekUsingTsInput = val
}

func (m *Mediafile) SetCopyTs(val bool) {
  m.copyTs = val
}

func (m *Mediafile) SetInputPath(val string) {
  m.inputPath = val
}

func (m *Mediafile) SetMuxDelay(val string) {
  m.muxDelay = val
}

func (m *Mediafile) SetOutputPath(val string) {
  m.outputPath = val
}

func (m *Mediafile) SetOutputFormat(val string) {
  m.outputFormat = val
}

func (m *Mediafile) SetNativeFramerateInput(val bool) {
  m.nativeFramerateInput = val
}

func (m *Mediafile) SetRtmpLive(val string) {
  m.rtmpLive = val
}

func (m *Mediafile) SetHlsSegmentDuration(val int) {
  m.hlsSegmentDuration = val
}

func (m *Mediafile) SetHlsPlaylistType(val string) {
  m.hlsPlaylistType = val
}

func (m *Mediafile) SetMetadata(v Metadata) {
  m.metadata = v
}

/*** GETTERS ***/

func (m *Mediafile) Filter() string {
  return m.filter
}

func (m *Mediafile) Aspect() string {
  return m.aspect
}

func (m *Mediafile) Resolution() string {
  return m.resolution
}

func (m *Mediafile) VideoBitrate() int {
  return m.videoBitRate
}

func (m *Mediafile) VideoBitRateTolerance() int {
  return m.videoBitRateTolerance
}

func (m *Mediafile) VideoMaxBitRate() int {
  return m.videoMaxBitRate
}

func (m *Mediafile) VideoMinBitRate() int {
  return m.videoMinBitrate
}

func (m *Mediafile) VideoCodec() string {
  return m.videoCodec
}

func (m *Mediafile) FrameRate() int {
  return m.frameRate
}

func (m *Mediafile) MaxKeyFrame() int {
  return m.maxKeyframe
}

func (m *Mediafile) MinKeyFrame() int {
  return m.minKeyframe
}

func (m *Mediafile) KeyFrameInterval() int {
  return m.keyframeInterval
}

func (m *Mediafile) AudioCodec() string {
  return m.audioCodec
}

func (m *Mediafile) AudioBitrate() int {
  return m.audioBitrate
}

func (m *Mediafile) AudioChannels() int {
  return m.audioChannels
}

func (m *Mediafile) BufferSize() int {
  return m.bufferSize
}

func (m *Mediafile) Threads() int {
  return m.threads
}

func (m *Mediafile) Target() string {
  return m.target
}

func (m *Mediafile) Duration() string {
  return m.duration
}

func (m *Mediafile) DurationInput() string {
  return m.durationInput
}

func (m *Mediafile) SeekTime() string {
  return m.seekTime
}

func (m *Mediafile) Preset() string {
  return m.preset
}

func (m *Mediafile) AudioProfile() string {
  return m.audioProfile
}

func (m *Mediafile) VideoProfile() string {
  return m.videoProfile
}

func (m *Mediafile) Tune() string {
  return m.tune
}

func (m *Mediafile) SeekTimeInput() string {
  return m.seekTimeInput
}

func (m *Mediafile) Quality() int {
  return m.quality
}

func (m *Mediafile) MuxDelay() string {
  return m.muxDelay
}

func (m *Mediafile) SeekUsingTsInput() bool {
  return m.seekUsingTsInput
}

func (m *Mediafile) CopyTs() bool {
  return m.copyTs
}

func (m *Mediafile) InputPath() string {
  return m.inputPath
}

func (m *Mediafile) OutputPath() string {
  return m.outputPath
}

func (m *Mediafile) OutputFormat() string {
  return m.outputFormat
}

func (m *Mediafile) NativeFramerateInput() bool {
  return m.nativeFramerateInput
}

func (m *Mediafile) RtmpLive() string {
  return m.rtmpLive
}

func (m *Mediafile) HlsSegmentDuration() int {
  return m.hlsSegmentDuration
}

func (m *Mediafile) HlsPlaylistType() string {
  return m.hlsPlaylistType
}

func (m *Mediafile) Metadata() Metadata {
  return m.metadata
}

/** OPTS **/
func (m *Mediafile) ToStrCommand() string {
  var strCommand string

  opts := [] string {
    "SeekTimeInput",
    "SeekUsingTsInput",
    "NativeFramerateInput",
    "DurationInput",
    "RtmpLive",
    "InputPath",

    "Aspect",
    "Resolution",
    "FrameRate",
    "VideoCodec",
    "VideoBitRate",
    "VideoBitRateTolerance",
    "VideoMaxBitRate",
    "VideoMinBitRate",
    "VideoProfile",
    "AudioCodec",
    "AudioBitRate",
    "AudioChannels",
    "AudioProfile",
    "Quality",
    "BufferSize",
    "MuxDelay",
    "Threads",
    "KeyframeInterval",
    "Preset",
    "Tune",
    "Target",
    "SeekTime",
    "Duration",
    "CopyTs",
    "OutputFormat",
    "HlsSegmentDuration",
    "HlsPlaylistType",
    "Filter",
    "OutputPath",
  }
  
  for _, name := range opts {
    opt :=  reflect.ValueOf(m).MethodByName(fmt.Sprintf("Obtain%s", name))
    if (opt != reflect.Value{}) {
      result := opt.Call([]reflect.Value{})

      if result[0].String() != "" {
        strCommand += " "
        strCommand += result[0].String()
      }
    }
  }

  return strCommand

}

func (m *Mediafile) ObtainFilter() string{
  if m.filter != "" {
    return fmt.Sprintf("-vf \"%s\"", m.filter)
  }
  return ""
}

func (m *Mediafile) ObtainAspect() string {
  // Set aspect
  if m.resolution != "" {
    resolution := strings.Split(m.resolution, "x")
    if len(resolution) != 0 {
      width, _ := strconv.ParseFloat(resolution[0], 64)
      height, _ := strconv.ParseFloat(resolution[1], 64)
      return fmt.Sprintf("-aspect %f", width/height)
    }
  }

  if m.aspect != "" {
    return fmt.Sprintf("-aspect %s", m.aspect)
  }
	return ""
}

func (m *Mediafile) ObtainInputPath() string {
  return fmt.Sprintf("-i \"%s\"", m.inputPath)
}

func (m *Mediafile) ObtainNativeFramerateInput() string {
  if m.nativeFramerateInput {
    return "-re"
  }
  return ""
}

func (m *Mediafile) ObtainOutputPath() string {
  return fmt.Sprintf("\"%s\"", m.outputPath)
}

func (m *Mediafile) ObtainVideoCodec() string {
  if m.videoCodec != "" {
    return fmt.Sprintf("-c:v %s", m.videoCodec)
  }
  return ""
}

func (m *Mediafile) ObtainFrameRate() string {
  if m.frameRate != 0 {
    return fmt.Sprintf("-r %d", m.frameRate)
  }
  return ""
}

func (m *Mediafile) ObtainResolution() string {
  if m.resolution != "" {
    return fmt.Sprintf("-s %s", m.resolution)
  }
  return ""
}

func (m *Mediafile) ObtainVideoBitRate() string {
  if m.videoBitRate != 0 {
    return fmt.Sprintf("-b:v %d", m.videoBitRate)
  }
  return ""
}

func (m *Mediafile) ObtainAudioCodec() string {
  if m.audioCodec != "" {
    return fmt.Sprintf("-c:a %s", m.audioCodec)
  }
  return ""
}

func (m *Mediafile) ObtainAudioBitRate() string {
  if m.audioBitrate != 0 {
    return fmt.Sprintf("-b:a %d", m.audioBitrate)
  }
  return ""
}

func (m *Mediafile) ObtainAudioChannels() string {
  if m.audioChannels != 0 {
    return fmt.Sprintf("-ac %d", m.audioChannels)
  }
  return ""
}

func (m *Mediafile) ObtainVideoMaxBitRate() string {
  if m.videoMaxBitRate != 0 {
    return fmt.Sprintf("-maxrate %dk", m.videoMaxBitRate)
  }
  return ""
}

func (m *Mediafile) ObtainVideoMinBitRate() string {
  if m.videoMinBitrate != 0 {
    return fmt.Sprintf("-minrate %dk", m.videoMinBitrate)
  }
  return ""
}

func (m *Mediafile) ObtainBufferSize() string {
  if m.bufferSize != 0 {
    return fmt.Sprintf("-bufsize %dk", m.bufferSize)
  }
  return ""
}

func (m *Mediafile) ObtainVideoBitRateTolerance() string {
  if m.videoBitRateTolerance != 0 {
    return fmt.Sprintf("-bt %dk", m.videoBitRateTolerance)
  }
  return ""
}

func (m *Mediafile) ObtainThreads() string {
  if m.threads != 0 {
    return fmt.Sprintf("-threads %d", m.threads)
  }
  return ""
}

func (m *Mediafile) ObtainTarget() string {
  if m.target != "" {
    return fmt.Sprintf("-target %s", m.target)
  }
  return ""
}

func (m *Mediafile) ObtainDuration() string {
  if m.duration != "" {
    return fmt.Sprintf("-t %s", m.duration)
  }
  return ""
}

func (m *Mediafile) ObtainDurationInput() string {
  if m.durationInput != "" {
    return fmt.Sprintf("-t %s", m.durationInput)
  }
  return ""
}

func (m *Mediafile) ObtainKeyframeInterval() string {
  if m.keyframeInterval != 0 {
    return fmt.Sprintf("-g %d", m.keyframeInterval)
  }
  return ""
}

func (m *Mediafile) ObtainSeekTime() string {
  if m.seekTime != "" {
    return fmt.Sprintf("-ss %s", m.seekTime)
  }
  return ""
}

func (m *Mediafile) ObtainSeekTimeInput() string {
  if m.seekTimeInput != "" {
    return fmt.Sprintf("-ss %s", m.seekTimeInput)
  }
  return ""
}

func (m *Mediafile) ObtainPreset() string {
  if m.preset != "" {
    return fmt.Sprintf("-preset %s", m.preset)
  }
  return ""
}

func (m *Mediafile) ObtainTune() string {
  if m.tune != "" {
    return fmt.Sprintf("-tune %s", m.tune)
  }
  return ""
}

func (m *Mediafile) ObtainQuality() string {
  if m.quality != 0 {
    return fmt.Sprintf("-crf %d", m.quality)
  }
  return ""
}

func (m *Mediafile) ObtainVideoProfile() string {
  if m.videoProfile != "" {
    return fmt.Sprintf("-profile:v %s", m.videoProfile)
  }
  return ""
}

func (m *Mediafile) ObtainAudioProfile() string {
  if m.audioProfile != "" {
    return fmt.Sprintf("-profile:a %s", m.audioProfile)
  }
  return ""
}

func (m *Mediafile) ObtainCopyTs() string {
  if m.copyTs {
    return "-copyts"
  }
  return ""
}

func (m *Mediafile) ObtainOutputFormat() string {
  if m.outputFormat != "" {
    return fmt.Sprintf("-f %s", m.outputFormat)
  }
  return ""
}

func (m *Mediafile) ObtainMuxDelay() string {
  if m.muxDelay != "" {
    return fmt.Sprintf("-muxdelay %s", m.muxDelay)
  }
  return ""
}

func (m *Mediafile) ObtainSeekUsingTsInput() string {
  if m.seekUsingTsInput {
    return "-seek_timestamp 1"
  }
	return ""
}

func (m *Mediafile) ObtainRtmpLive() string {
  if m.rtmpLive != "" {
    return fmt.Sprintf("-rtmp_live %s", m.rtmpLive)
  } else {
    return ""
  }
}

func (m *Mediafile) ObtainHlsPlaylistType() string {
  if m.hlsPlaylistType != "" {
    return fmt.Sprintf("-hls_playlist_type %s", m.hlsPlaylistType)
  } else {
    return ""
  }
}

func (m *Mediafile) ObtainHlsSegmentDuration() string {
  if m.hlsSegmentDuration != 0 {
    return fmt.Sprintf("-hls_time %d", m.hlsSegmentDuration)
  } else {
    return ""
  }
}
