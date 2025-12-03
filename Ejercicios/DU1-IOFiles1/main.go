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
	ID, Nombre, Categoria string
	Precio                float64
	Stock                 int
}

type Transaccion struct {
	Tipo, IDProducto, Fecha string
	Cantidad                int
}

func main() {
	const (
		inventarioArchivo    = "inventario.txt"
		transaccionesArchivo = "transacciones.txt"
		inventarioOut        = "inventario_actualizado.txt"
		reporteBajoStock     = "productos_bajo_stock.txt"
		logErrores           = "errores.log"
		limiteBajoStock      = 10
	)

	// Leer datos
	productos, err := leerInventario(inventarioArchivo)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	transacciones, err := leerTransacciones(transaccionesArchivo)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Procesar y obtener errores
	errores := procesarTransacciones(productos, transacciones)

	// Guardar resultados
	_ = escribirInventario(productos, inventarioOut)
	_ = generarReporteBajoStock(productos, limiteBajoStock, reporteBajoStock)
	_ = escribirLog(errores, logErrores)
}

// ---------------- Funciones ----------------

// Leer archivo genérico (devuelve líneas sin cabecera)
func leerArchivo(nombre string) ([]string, error) {
	file, err := os.Open(nombre)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if !scanner.Scan() {
		return nil, fmt.Errorf("archivo vacío: %s", nombre)
	}

	var lineas []string
	for scanner.Scan() {
		linea := strings.TrimSpace(scanner.Text())
		if linea != "" {
			lineas = append(lineas, linea)
		}
	}
	return lineas, scanner.Err()
}

// Leer inventario
func leerInventario(nombre string) (map[string]*Producto, error) {
	lineas, err := leerArchivo(nombre)
	if err != nil {
		return nil, err
	}

	productos := make(map[string]*Producto)
	for i, l := range lineas {
		campos := strings.Split(l, ",")
		if len(campos) != 5 {
			fmt.Printf("Línea %d inválida en inventario\n", i+2)
			continue
		}
		precio, err1 := strconv.ParseFloat(campos[3], 64)
		stock, err2 := strconv.Atoi(campos[4])
		if err1 != nil || err2 != nil {
			fmt.Printf("Línea %d: error en precio o stock\n", i+2)
			continue
		}
		productos[campos[0]] = &Producto{
			ID: campos[0], Nombre: campos[1], Categoria: campos[2],
			Precio: precio, Stock: stock,
		}
	}
	return productos, nil
}

// Leer transacciones
func leerTransacciones(nombre string) ([]Transaccion, error) {
	lineas, err := leerArchivo(nombre)
	if err != nil {
		return nil, err
	}

	var transacciones []Transaccion
	for i, l := range lineas {
		campos := strings.Split(l, ",")
		if len(campos) != 4 {
			fmt.Printf("Línea %d inválida en transacciones\n", i+2)
			continue
		}
		cant, err := strconv.Atoi(campos[2])
		if err != nil {
			fmt.Printf("Línea %d: cantidad inválida\n", i+2)
			continue
		}
		transacciones = append(transacciones, Transaccion{
			Tipo: campos[0], IDProducto: campos[1], Cantidad: cant, Fecha: campos[3],
		})
	}
	return transacciones, nil
}

// Procesar transacciones
func procesarTransacciones(productos map[string]*Producto, trans []Transaccion) []string {
	var errores []string

	for _, t := range trans {
		p, existe := productos[t.IDProducto]
		if !existe {
			errores = append(errores, fmt.Sprintf("[%s] ERROR: Producto %s no encontrado en transacción de tipo %s",
				t.Fecha, t.IDProducto, t.Tipo))
			continue
		}

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
			errores = append(errores, fmt.Sprintf("[%s] ERROR: Tipo de transacción desconocido: %s",
				t.Fecha, t.Tipo))
		}
	}
	return errores
}

// Escribir inventario actualizado
func escribirInventario(productos map[string]*Producto, nombre string) error {
	file, err := os.Create(nombre)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	fmt.Fprintln(writer, "ID,Nombre,Categoria,Precio,Stock")
	for _, p := range productos {
		fmt.Fprintf(writer, "%s,%s,%s,%.2f,%d\n", p.ID, p.Nombre, p.Categoria, p.Precio, p.Stock)
	}
	return writer.Flush()
}

// Reporte de bajo stock
func generarReporteBajoStock(productos map[string]*Producto, limite int, nombre string) error {
	file, err := os.Create(nombre)
	if err != nil {
		return err
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

// Escribir log de errores
func escribirLog(errores []string, nombre string) error {
	if len(errores) == 0 {
		return nil
	}
	file, err := os.Create(nombre)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, e := range errores {
		timestamp := time.Now().Format("2006-01-02 15:04:05")
		fmt.Fprintf(writer, "[%s] %s\n", timestamp, e)
	}
	return writer.Flush()
}
