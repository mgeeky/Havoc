package logr

import (
    "fmt"
    "log"
    "os"

    "Havoc/pkg/common"
    "Havoc/pkg/logger"
)

func (l Logr) AddAgentInput(AgentType, AgentID, User, TaskID, Input string, time string) {
    var (
        DemonPath    = l.AgentPath + "/" + AgentID
        DemonLogFile = DemonPath + "/Console_" + AgentID + ".log"
        InputString  string
    )

    if _, err := os.Stat(DemonPath); os.IsNotExist(err) {
        if err = os.Mkdir(DemonPath, os.ModePerm); err != nil {
            logger.Error("Failed to create Logr demon " + AgentID + " folder: " + err.Error())
            return
        }
    }

    f, err := os.OpenFile(DemonLogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal(err)
    }

    InputString = fmt.Sprintf("\n[%v] [%v] %v => %v\n", time, User, AgentType, Input)

    _, err = f.Write([]byte(InputString))
    if err != nil {
        logger.Error("Failed to write to File [" + DemonLogFile + "]: " + err.Error())
        return
    }
}

func (l Logr) AddAgentRaw(AgentID, Raw string) {
    var (
        DemonPath    = l.AgentPath + "/" + AgentID
        DemonLogFile = DemonPath + "/Console_" + AgentID + ".log"
    )

    if _, err := os.Stat(DemonPath); os.IsNotExist(err) {
        if err = os.Mkdir(DemonPath, os.ModePerm); err != nil {
            logger.Error("Failed to create Logr demon " + AgentID + " folder: " + err.Error())
            return
        }
    }

    f, err := os.OpenFile(DemonLogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal(err)
    }

    _, err = f.Write([]byte(Raw))
    if err != nil {
        logger.Error("Failed to write to File [" + DemonLogFile + "]: " + err.Error())
        return
    }
}

func (l Logr) DemonAddOutput(DemonID string, Output map[string]string, time string) {
    var (
        DemonPath    = l.AgentPath + "/" + DemonID
        DemonLogFile = DemonPath + "/Console_" + DemonID + ".log"
    )

    if _, err := os.Stat(DemonPath); os.IsNotExist(err) {
        if err = os.Mkdir(DemonPath, os.ModePerm); err != nil {
            logger.Error("Failed to create Logr demon " + DemonID + " folder: " + err.Error())
            return
        }
    }

    f, err := os.OpenFile(DemonLogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal(err)
    }

    var OutputString string

    if len(Output["Message"]) > 0 {

        if Output["Type"] == "Good" {
            OutputString = fmt.Sprintf("[%v] [+] %v\n", time, Output["Message"])
        } else if Output["Type"] == "Error" {
            OutputString = fmt.Sprintf("[%v] [-] %v\n", time, Output["Message"])
        } else if Output["Type"] == "Info" {
            OutputString = fmt.Sprintf("[%v] [*] %v\n", time, Output["Message"])
        } else {
            OutputString = fmt.Sprintf("[%v] [^] %v\n", time, Output["Message"])
        }

    }

    if len(Output["Output"]) > 0 {
        OutputString += Output["Output"]
    }

    _, err = f.Write([]byte(OutputString))
    if err != nil {
        logger.Error("Failed to write to File [" + DemonLogFile + "]: " + err.Error())
        return
    }
}

func (l Logr) DemonAddDownloadedFile(DemonID, FileName string, FileBytes []byte) {
    var (
        DemonPath        = l.AgentPath + "/" + DemonID
        DemonDownloadDir = DemonPath + "/Download"
        DemonDownload    = DemonDownloadDir + "/" + FileName
    )

    if _, err := os.Stat(DemonPath); os.IsNotExist(err) {
        if err = os.Mkdir(DemonPath, os.ModePerm); err != nil {
            logger.Error("Failed to create Logr demon " + DemonID + " folder: " + err.Error())
            return
        }
    }

    if _, err := os.Stat(DemonDownloadDir); os.IsNotExist(err) {
        if err = os.Mkdir(DemonDownloadDir, os.ModePerm); err != nil {
            logger.Error("Failed to create Logr demon " + DemonID + " screenshot folder: " + err.Error())
            return
        }
    }

    f, err := os.Create(DemonDownload)
    if err != nil {
        logger.Error("Failed to create file: " + err.Error())
        return
    }

    defer f.Close()

    _, err = f.Write(FileBytes)
    if err != nil {
        logger.Error("Failed to write png file: " + err.Error())
        return
    }
}

func (l Logr) DemonSaveScreenshot(DemonID, Name string, BmpBytes []byte) {
    var (
        DemonPath          = l.AgentPath + "/" + DemonID
        DemonScreenshotDir = DemonPath + "/Screenshots"
        DemonScreenshot    = DemonScreenshotDir + "/" + Name
    )

    if _, err := os.Stat(DemonPath); os.IsNotExist(err) {
        if err = os.Mkdir(DemonPath, os.ModePerm); err != nil {
            logger.Error("Failed to create Logr demon " + DemonID + " folder: " + err.Error())
            return
        }
    }

    if _, err := os.Stat(DemonScreenshotDir); os.IsNotExist(err) {
        if err = os.Mkdir(DemonScreenshotDir, os.ModePerm); err != nil {
            logger.Error("Failed to create Logr demon " + DemonID + " screenshot folder: " + err.Error())
            return
        }
    }

    f, err := os.Create(DemonScreenshot)
    if err != nil {
        logger.Error("Failed to create file: " + err.Error())
        return
    }
    defer f.Close()

    _, err = f.Write(common.Bmp2Png(BmpBytes))
    if err != nil {
        logger.Error("Failed to write png file: " + err.Error())
        return
    }
}