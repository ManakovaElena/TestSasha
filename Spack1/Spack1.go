package Spack1

import(
	"log/slog"
)

func Nop(){
	slog.Info("Spack create", slog.String("S", "A"))
}