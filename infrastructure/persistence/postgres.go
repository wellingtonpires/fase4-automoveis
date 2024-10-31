package persistence

import (
	"database/sql"
	"fmt"
	"sort"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	_ "github.com/lib/pq"
	"github.com/wellingtonpires/fase4-automoveis/domain/entity/veiculo"
)

const conexaoAberta = "Conexao aberta com o banco!"

func OpenConnection() (*sql.DB, error) {
	db, err := sql.Open("postgres", "postgresql://postgres:123@postgres:5432/postgres?sslmode=disable")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Conectado ao banco!")
	}

	err = db.Ping()

	return db, err
}

func ConsultaPorId(id int) (veiculos veiculo.Veiculo) {
	con, err := OpenConnection()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(conexaoAberta)
	}
	defer con.Close()
	rows, err := con.Query(`SELECT * FROM veiculos WHERE id = $1`, id)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var v veiculo.Veiculo
		rows.Scan(&v.Marca, &v.Modelo, &v.Ano, &v.Cor, &v.Preco, &v.Flagvendido, &v.Id)
	}

	return veiculos
}

func extraiCpf(tokenJwt string) string {
	var cpf string
	claims := jwt.MapClaims{}
	_, _, err := jwt.NewParser().ParseUnverified(tokenJwt, claims)
	if err != nil {
		fmt.Println(err.Error())
	}
	for key, val := range claims {
		cpf = fmt.Sprintf("%v,%v", key, val)
	}

	return strings.Split(cpf, ",")[1]
}

func Checkout(v veiculo.Veiculo, cpfvenda string) {
	con, err := OpenConnection()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(conexaoAberta)
	}

	defer con.Close()
	_, err = con.Exec(`UPDATE veiculos SET flagvendido = $1, cpf = $2 WHERE id = $3`, true, extraiCpf(cpfvenda), v.Id)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func ConsultaPorPreco() (veiculos []veiculo.Veiculo) {
	con, err := OpenConnection()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(conexaoAberta)
	}
	defer con.Close()
	rows, err := con.Query(`SELECT * FROM veiculos WHERE flagvendido = false`)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var v veiculo.Veiculo
		rows.Scan(&v.Marca, &v.Modelo, &v.Ano, &v.Cor, &v.Preco, &v.Flagvendido, &v.Id)
		veiculos = append(veiculos, v)
	}
	sort.Slice(veiculos, func(i, j int) bool {
		return veiculos[i].Preco < veiculos[j].Preco
	})
	return veiculos
}

func ConsultaVendidos() (veiculos []veiculo.Veiculo) {
	con, err := OpenConnection()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(conexaoAberta)
	}
	defer con.Close()
	flag := "y"
	rows, err := con.Query(`SELECT * FROM veiculos WHERE flagvendido = $1`, flag)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var v veiculo.Veiculo
		rows.Scan(&v.Marca, &v.Modelo, &v.Ano, &v.Cor, &v.Preco, &v.Flagvendido, &v.Id)
		veiculos = append(veiculos, v)
	}
	sort.Slice(veiculos, func(i, j int) bool {
		return veiculos[i].Preco < veiculos[j].Preco
	})
	return veiculos
}

func CadastraVeiculo(v veiculo.Veiculo) {
	con, err := OpenConnection()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(conexaoAberta)
	}
	defer con.Close()
	_, err = con.Exec(`INSERT INTO veiculos VALUES ($1, $2, $3, $4, $5, $6)`, v.Marca, v.Modelo, v.Ano, v.Cor, v.Preco, v.Flagvendido)
	if err != nil {
		fmt.Println(err.Error())
	}
}

// Caso um campo esteja vazio irá fazer update dele vazio na base. Precisa corrigir.
func AtualizaVeiculo(v veiculo.Veiculo) {
	con, err := OpenConnection()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(conexaoAberta)
	}
	defer con.Close()
	_, err = con.Exec(`UPDATE veiculos SET marca = $1, modelo = $2, ano = $3, cor = $4, preco = $5, flagvendido = $6 WHERE id = $7`, v.Marca, v.Modelo, v.Ano, v.Cor, v.Preco, v.Flagvendido, v.Id)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func ExcluiVeiculo(v veiculo.Veiculo) {
	con, err := OpenConnection()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(conexaoAberta)
	}
	defer con.Close()
	_, err = con.Exec(`DELETE FROM veiculos WHERE id = $1`, v.Id)
	if err != nil {
		fmt.Println(err.Error())
	}
}
