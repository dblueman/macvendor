package macvendor

import (
   "encoding/json"
   "os"
   "sync"

   "github.com/umahmood/macvendors"
)

type MacVendor struct {
   lowlevel *macvendors.API
   lock sync.Mutex
   cache map[string]string
}

var (
   cachePath string
   nameMap = map[string]string{
      "Apple, Inc.":                         "Apple",
      "Microsoft Corporation":               "Microsoft",
      "No vendor":                           "Unknown",
      "Xiaomi Communications Co Ltd":        "Xiaomi",
      "SAMSUNG ELECTRO-MECHANICS(THAILAND)": "Samsung",
      "Intel Corporate":                     "Intel",
      "HUAWEI TECHNOLOGIES CO.,LTD":         "Huawei",
   }
)

func (m *MacVendor) Lookup(mac string) (string, error) {
   prefix := mac[:8]

   m.lock.Lock()
   defer m.lock.Unlock()

   vendor, ok := m.cache[prefix]
   if !ok {
      name, err := m.lowlevel.Name(prefix)
      if err != nil {
         return "", err
      }

      m.cache[prefix] = name
      serialised, err := json.Marshal(m.cache)
      if err != nil {
         return "", err
      }

      err = os.WriteFile(cachePath, serialised, 0644)
      if err != nil {
         panic(err)
         return "", err
      }

      vendor = name
   }

   vendor2, ok := nameMap[vendor]
   if !ok {
      vendor2 = vendor
   }

   return vendor2, nil
}

func New() (*MacVendor, error) {
   cacheDir, err := os.UserCacheDir()
   if err != nil {
      return nil, err
   }

   cachePath = cacheDir + "/macvendor.json"
   inst := MacVendor{lowlevel: macvendors.New()}

   serialised, err := os.ReadFile(cachePath)
   // FIXME check for os.IsNotExists error
   if err == nil {
      err = json.Unmarshal(serialised, &inst.cache)
      if err != nil {
         return nil, err
      }
   } else {
      inst.cache = map[string]string{}
   }

   return &inst, nil
}
