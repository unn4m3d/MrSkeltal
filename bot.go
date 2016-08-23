package main

import (
    "log"
    "strings"
    "time"
    "github.com/go-telegram-bot-api/telegram-bot-api"
    "math/rand"
    "strconv"
)

func main() {
    bot, err := tgbotapi.NewBotAPI("TOKEN")
    if err != nil {
        log.Panic(err)
    }

    // bot.Debug = true

    log.Printf("Authorized on account %s", bot.Self.UserName)

    u := tgbotapi.NewUpdate(0)
    u.Timeout = 60

    updates, err := bot.GetUpdatesChan(u)

    for update := range updates {
        if update.Message == nil {
            continue
        }

        log.Printf("<--[%s] %s", update.Message.From.UserName, update.Message.Text)
        if strings.ToLower(update.Message.Text) == "спасибо, мистер дудец" ||
        strings.ToLower(update.Message.Text) == "спасибо мистер дудец"{
          rand.Seed(time.Now().UTC().UnixNano())
          number := rand.Intn(17)
          log.Printf(strconv.Itoa(number))
          if number == 0 || number == 1 || number == 2 || number == 3 || number == 4 || number == 5 {
            msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.From.FirstName+posPhrases[number])
            bot.Send(msg)
            log.Printf("-->[%s] %s", update.Message.From.UserName, msg.Text)
          } else if number == 6 ||number == 7 ||number == 8 ||number == 9{
            i, _ := strconv.Atoi(posPhrases[number])
              msg := tgbotapi.NewPhotoShare(update.Message.Chat.ID, sharedImages[i])
              message, err :=  bot.Send(msg)
              checkErr(err, "Sending failed")
              log.Printf("-->[%s] %s", update.Message.From.UserName, message.Photo)
          } else if number == 10 ||number == 11 ||number == 12 ||number == 13{
            i, _ := strconv.Atoi(posPhrases[number])
            msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.From.FirstName+videos[i])
            bot.Send(msg)
            log.Printf("-->[%s] %s", update.Message.From.UserName, msg.Text)
          } else if number == 14 ||number == 15 ||number == 16 ||number == 17{
            i, _ := strconv.Atoi(posPhrases[number])
              msg := tgbotapi.NewDocumentShare(update.Message.Chat.ID, sharedGifs[i])
              message, err :=  bot.Send(msg)
              checkErr(err, "Sending failed")
              log.Printf("-->[%s] %s", update.Message.From.UserName, gifs[i]+" " + message.Document.FileID)
          }
        } else if strings.Contains(strings.ToLower(update.Message.Text), "нахуй") ||
        strings.Contains(strings.ToLower(update.Message.Text), "пизда") ||
        strings.Contains(strings.ToLower(update.Message.Text), "уебывай") ||
        strings.Contains(strings.ToLower(update.Message.Text), "пидор") ||
        strings.Contains(strings.ToLower(update.Message.Text), "гей") ||
        strings.Contains(strings.ToLower(update.Message.Text), "хуй") ||
        strings.Contains(strings.ToLower(update.Message.Text), "сосать") ||
        strings.Contains(strings.ToLower(update.Message.Text), "давалка") ||
        strings.Contains(strings.ToLower(update.Message.Text), "уебок"){
          rand.Seed(time.Now().UTC().UnixNano())
          number := rand.Intn(9)
          msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.From.FirstName+negPhrases[number])
          bot.Send(msg)
          log.Printf("-->[%s] %s", update.Message.From.UserName, msg.Text)
        } else if strings.Contains(strings.ToLower(update.Message.Text), "чем ты лучше"){
          msg := tgbotapi.NewPhotoShare(update.Message.Chat.ID, "AgADAgADtacxG9RzuQ5j3ZDJp3kLqjkncQ0ABFSkkl0_Vpx4O3ABAAEC")
          message, err :=  bot.Send(msg)
          checkErr(err, "Sending failed")
          log.Printf("-->[%s] %s", update.Message.From.UserName, message.Photo)
        }
    }
}
