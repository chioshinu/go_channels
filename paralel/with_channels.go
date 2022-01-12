// package main
 
// import (
//  "bytes"
//  "crypto/md5"
//  "fmt"
//  "net/http"
//  "net/url"
//  "os"
//  "os/signal"
//  "strconv"
//  "strings"
//  "sync"
//  "syscall"
//  "time"
 
//  "github.com/cheggaaa/pb"
//  "source.masterofcode.com/vitalii.hurin/tagible-for-travel-go-backend/config"
//  "source.masterofcode.com/vitalii.hurin/tagible-for-travel-go-backend/entity"
//  "source.masterofcode.com/vitalii.hurin/tagible-for-travel-go-backend/repository"
// )
 
// func main() {
//  config.CheckDB()
//  config.CheckImageStorageDomain()
 
//  if len(os.Args) < 3 {
//    fmt.Printf("USAGE: tft-cli-pano-thumbnail-checker <start-id> <finish-id>\n")
//    os.Exit(1)
//  }
 
//  startID, _ := strconv.ParseInt(os.Args[1], 10, 64)
//  finishID, _ := strconv.ParseInt(os.Args[2], 10, 64)
//  panosCount := int(finishID - startID + 1)
//  if panosCount <= 0 {
//    fmt.Printf("Bad range of scanning IDs: [%v..%v]\n", startID, finishID)
//    os.Exit(1)
//  }
 
//  workers := int(*config.WorkersCount)
//  if panosCount < workers {
//    workers = panosCount
//  }
 
//  panos := repository.NewPano(config.DB, false)
 
//  c := make(chan os.Signal, 1)
//  signal.Notify(c, syscall.SIGINT)
 
//  begin := time.Now()
//  fmt.Printf("Pano thumbnails checking starting at %v. Total words count: %v\n", begin, panosCount)
 
//  workChan := make(chan int64, panosCount)
//  for i := startID; i <= finishID; i++ {
//    workChan <- i
//  }
 
//  bar := pb.New(panosCount)
//  bar.Start()
 
//  var wg sync.WaitGroup
//  wg.Add(workers)
//  for i := 0; i < workers; i++ {
//    go func(workerID int, bar *pb.ProgressBar) {
//      defer wg.Done()
//      for {
//        if len(c) != 0 {
//          return
//        }
//        select {
//        case w := <-workChan:
//          p, err := panos.FindByID(nil, entity.PanoID(w))
//          if err != nil || p.Type == entity.PanoTypeMap {
//            bar.Increment()
//            continue
//          }
//          idString := strconv.FormatInt(int64(p.ID), 10)
 
//          var filename bytes.Buffer
//          filename.WriteString(*config.ImageStorageDomain)
//          filename.WriteString("thumbnails/pano/")
//          filename.WriteString(idString)
//          filename.WriteByte('-')
//          filename.WriteString(fmt.Sprintf("%x", md5.Sum([]byte(p.PanoID))))
 
//          res, err := http.Get(filename.String() + "_big.png")
//          if err == nil {
//            res.Body.Close()
//          }
//          if err != nil || res.StatusCode < 200 || res.StatusCode >= 300 || res.ContentLength < 300 {
//            err = send(&url.Values{"panoid": {idString}}, http.DefaultClient)
//            if err != nil {
//              fmt.Printf("Worker %v. Panoid: %s. %v\n", workerID, idString, err.Error())
//              workChan <- w
//              continue
//            }
//            bar.Increment()
//            continue
//          }
//          res, err = http.Get(filename.String() + "_small.png")
//          if err == nil {
//            res.Body.Close()
//          }
//          if err != nil || res.StatusCode < 200 || res.StatusCode >= 300 || res.ContentLength < 300 {
//            err = send(&url.Values{"panoid": {idString}}, http.DefaultClient)
//            if err != nil {
//              fmt.Printf("Worker %v. Panoid: %s. %v\n", workerID, idString, err.Error())
//              workChan <- w
//              continue
//            }
//            bar.Increment()
//            continue
//          }
//          bar.Increment()
//        default:
//          return
//        }
//      }
//    }(i, bar)
//  }
//  wg.Wait()
//  bar.Finish()
//  fmt.Printf("Pano thumbnails checking finished at %v. Took: %v\n", time.Now(), time.Since(begin))
// }
 
// func send(v *url.Values, c *http.Client) error {
//  r, err := http.NewRequest("POST", "http://localhost:8484/rasterize", strings.NewReader(v.Encode()))
//  if err != nil {
//    return fmt.Errorf("Failed to send page to page processing service: %v", err)
//  }
//  r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
//  postResp, err := c.Do(r)
//  if err != nil {
//    return fmt.Errorf("Failed to send page to page processing service: %v", err)
//  }
//  postResp.Body.Close()
//  if postResp.StatusCode < 200 || postResp.StatusCode >= 300 {
//    return fmt.Errorf("Bad response from page processing service: %s", postResp.Status)
//  }
//  return nil
// }
 
 

