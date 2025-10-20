package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Producto struct {
	ID        string
	Nombre    string
	Categoria string
	Precio    float64
	Stock     int
}

type Transaccion struct {
	Tipo       string
	IDProducto string
	Cantidad   int
	Fecha      string
}

func main() {
	inventarioArchivo := "inventario.txt"
	transaccionesArchivo := "transacciones.txt"

	// Leer inventario
	productos, err := leerInventario(inventarioArchivo)
	if err != nil {
		fmt.Println("Error al leer inventario:", err)
		return
	}

	// Leer transacciones
	transacciones, err := leerTransacciones(transaccionesArchivo)
	if err != nil {
		fmt.Println("Error al leer transacciones:", err)
		return
	}

	// Procesar transacciones
	errores := procesarTransacciones(productos, transacciones)

	// Escribir inventario actualizado
	if err := escribirInventario(productos, "inventario_actualizado.txt"); err != nil {
		fmt.Println("Error al escribir inventario actualizado:", err)
	}

	// Generar reporte de bajo stock (<10 unidades)
	if err := generarReporteBajoStock(productos, 10); err != nil {
		fmt.Println("Error al generar reporte de bajo stock:", err)
	}

	// Escribir errores en log
	if err := escribirLog(errores, "errores.log"); err != nil {
		fmt.Println("Error al escribir log de errores:", err)
	}

}

// ---------------------- Funciones ----------------------

// 1. Leer inventario
func leerInventario(nombreArchivo string) ([]Producto, error) {
	file, err := os.Open(nombreArchivo)
	if err != nil {
		return nil, fmt.Errorf("no se pudo abrir inventario: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	productos := []Producto{}

	// Saltar cabecera
	if !scanner.Scan() {
		return nil, fmt.Errorf("archivo de inventario vacío")
	}

	lineaNum := 1
	for scanner.Scan() {
		lineaNum++
		linea := strings.TrimSpace(scanner.Text())
		if linea == "" {
			continue
		}
		campos := strings.Split(linea, ",")
		if len(campos) != 5 {
			fmt.Printf("Línea %d ignorada: formato incorrecto\n", lineaNum)
			continue
		}

		precio, err := strconv.ParseFloat(campos[3], 64)
		if err != nil {
			fmt.Printf("Línea %d: precio inválido\n", lineaNum)
			continue
		}
		stock, err := strconv.Atoi(campos[4])
		if err != nil {
			fmt.Printf("Línea %d: stock inválido\n", lineaNum)
			continue
		}

		productos = append(productos, Producto{
			ID:        campos[0],
			Nombre:    campos[1],
			Categoria: campos[2],
			Precio:    precio,
			Stock:     stock,
		})
	}
	return productos, nil
}

// 2. Leer transacciones
func leerTransacciones(nombreArchivo string) ([]Transaccion, error) {
	file, err := os.Open(nombreArchivo)
	if err != nil {
		return nil, fmt.Errorf("no se pudo abrir transacciones: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	transacciones := []Transaccion{}

	// Saltar cabecera
	if !scanner.Scan() {
		return nil, fmt.Errorf("archivo de transacciones vacío")
	}

	lineaNum := 1
	for scanner.Scan() {
		lineaNum++
		linea := strings.TrimSpace(scanner.Text())
		if linea == "" {
			continue
		}
		campos := strings.Split(linea, ",")
		if len(campos) != 4 {
			fmt.Printf("⚠️ Línea %d ignorada: formato incorrecto\n", lineaNum)
			continue
		}

		cantidad, err := strconv.Atoi(campos[2])
		if err != nil {
			fmt.Printf("⚠️ Línea %d: cantidad inválida\n", lineaNum)
			continue
		}

		transacciones = append(transacciones, Transaccion{
			Tipo:       campos[0],
			IDProducto: campos[1],
			Cantidad:   cantidad,
			Fecha:      campos[3],
		})
	}
	return transacciones, nil
}

// 3. Procesar transacciones
func procesarTransacciones(productos []Producto, transacciones []Transaccion) []string {
	errores := []string{}

	for _, t := range transacciones {
		// Buscar producto
		idx := -1
		for i, p := range productos {
			if p.ID == t.IDProducto {
				idx = i
				break
			}
		}
		if idx == -1 {
			errores = append(errores, fmt.Sprintf("[%s] ERROR: Producto %s no encontrado en transacción de tipo %s", t.Fecha, t.IDProducto, t.Tipo))
			continue
		}

		p := &productos[idx]

		switch strings.ToUpper(t.Tipo) {
		case "VENTA":
			if p.Stock < t.Cantidad {
				errores = append(errores, fmt.Sprintf("[%s] ERROR: Stock insuficiente para venta. Producto: %s, Stock actual: %d, Cantidad solicitada: %d",
					t.Fecha, p.ID, p.Stock, t.Cantidad))
				continue
			}
			p.Stock -= t.Cantidad
		case "COMPRA", "DEVOLUCION":
			p.Stock += t.Cantidad
		default:
			errores = append(errores, fmt.Sprintf("[%s] ERROR: Tipo de transacción desconocido: %s", t.Fecha, t.Tipo))
		}
	}

	return errores
}

// 4. Escribir inventario actualizado
func escribirInventario(productos []Producto, nombreArchivo string) error {
	file, err := os.Create(nombreArchivo)
	if err != nil {
		return fmt.Errorf("no se pudo crear archivo %s: %w", nombreArchivo, err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	fmt.Fprintln(writer, "ID,Nombre,Categoria,Precio,Stock")
	for _, p := range productos {
		fmt.Fprintf(writer, "%s,%s,%s,%.2f,%d\n", p.ID, p.Nombre, p.Categoria, p.Precio, p.Stock)
	}
	return writer.Flush()
}

// 5. Generar reporte de bajo stock
func generarReporteBajoStock(productos []Producto, limite int) error {
	file, err := os.Create("productos_bajo_stock.txt")
	if err != nil {
		return fmt.Errorf("no se pudo crear productos_bajo_stock.txt: %w", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	fmt.Fprintln(writer, "ALERTA: PRODUCTOS CON BAJO STOCK")
	fmt.Fprintln(writer, "================================")

	contador := 0
	for _, p := range productos {
		if p.Stock < limite {
			fmt.Fprintf(writer, "ID: %s | %s | Stock actual: %d unidades\n", p.ID, p.Nombre, p.Stock)
			contador++
		}
	}
	fmt.Fprintf(writer, "Total de productos con bajo stock: %d\n", contador)
	return writer.Flush()
}

// 6. Escribir log de errores
func escribirLog(errores []string, nombreArchivo string) error {
	if len(errores) == 0 {
		return nil
	}

	file, err := os.Create(nombreArchivo)
	if err != nil {
		return fmt.Errorf("no se pudo crear log de errores: %w", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, e := range errores {
		// Añadir timestamp real al log
		timestamp := time.Now().Format("2006-01-02 15:04:05")
		fmt.Fprintf(writer, "[%s] %s\n", timestamp, e)
	}
	return writer.Flush()
}
