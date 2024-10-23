package spack2
import
    ("log/slog")

func Nop() {
    slog.Info("Here package nubmer 2",
    slog.String("Package", "Spack2"))
}