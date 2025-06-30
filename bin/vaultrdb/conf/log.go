package conf

import "github.com/jnnkrdb/vaultrdb/pkg/logging"

var Log = logging.GetLogger(YC.Log.FormatJSON, YC.Log.Level)
