package sql

import (
	"fmt"
	"strconv"
	"strings"
)

const CommaSpace = ", "

type Mysql struct {
	Parts []string
}

func (qb *Mysql) Select(fields ...string) SqlBuilder {
	qb.Parts = append(qb.Parts, "SELECT", strings.Join(fields, CommaSpace))
	return qb
}

func (qb *Mysql) ForUpdate() SqlBuilder {
	qb.Parts = append(qb.Parts, "FOR UPDATE")
	return qb
}

func (qb *Mysql) From(tables ...string) SqlBuilder {
	qb.Parts = append(qb.Parts, "FROM", strings.Join(tables, CommaSpace))
	return qb
}

func (qb *Mysql) InnerJoin(table string) SqlBuilder {
	qb.Parts = append(qb.Parts, "INNER JOIN", table)
	return qb
}

func (qb *Mysql) LeftJoin(table string) SqlBuilder {
	qb.Parts = append(qb.Parts, "LEFT JOIN", table)
	return qb
}

func (qb *Mysql) RightJoin(table string) SqlBuilder {
	qb.Parts = append(qb.Parts, "RIGHT JOIN", table)
	return qb
}

func (qb *Mysql) On(cond string) SqlBuilder {
	qb.Parts = append(qb.Parts, "ON", cond)
	return qb
}

func (qb *Mysql) Where(cond string) SqlBuilder {
	qb.Parts = append(qb.Parts, "WHERE", cond)
	return qb
}

func (qb *Mysql) And(cond string) SqlBuilder {
	qb.Parts = append(qb.Parts, "AND", cond)
	return qb
}

func (qb *Mysql) Or(cond string) SqlBuilder {
	qb.Parts = append(qb.Parts, "OR", cond)
	return qb
}

func (qb *Mysql) In(vals ...string) SqlBuilder {
	qb.Parts = append(qb.Parts, "IN", "(", strings.Join(vals, CommaSpace), ")")
	return qb
}

func (qb *Mysql) OrderBy(fields ...string) SqlBuilder {
	qb.Parts = append(qb.Parts, "ORDER BY", strings.Join(fields, CommaSpace))
	return qb
}

func (qb *Mysql) Asc() SqlBuilder {
	qb.Parts = append(qb.Parts, "ASC")
	return qb
}

func (qb *Mysql) Desc() SqlBuilder {
	qb.Parts = append(qb.Parts, "DESC")
	return qb
}

func (qb *Mysql) Limit(limit int) SqlBuilder {
	qb.Parts = append(qb.Parts, "LIMIT", strconv.Itoa(limit))
	return qb
}

func (qb *Mysql) Offset(offset int) SqlBuilder {
	qb.Parts = append(qb.Parts, "OFFSET", strconv.Itoa(offset))
	return qb
}

func (qb *Mysql) GroupBy(fields ...string) SqlBuilder {
	qb.Parts = append(qb.Parts, "GROUP BY", strings.Join(fields, CommaSpace))
	return qb
}

func (qb *Mysql) Having(cond string) SqlBuilder {
	qb.Parts = append(qb.Parts, "HAVING", cond)
	return qb
}

func (qb *Mysql) Update(tables ...string) SqlBuilder {
	qb.Parts = append(qb.Parts, "UPDATE", strings.Join(tables, CommaSpace))
	return qb
}

func (qb *Mysql) Set(kv ...string) SqlBuilder {
	qb.Parts = append(qb.Parts, "SET", strings.Join(kv, CommaSpace))
	return qb
}

func (qb *Mysql) Delete(tables ...string) SqlBuilder {
	qb.Parts = append(qb.Parts, "DELETE")
	if len(tables) != 0 {
		qb.Parts = append(qb.Parts, strings.Join(tables, CommaSpace))
	}
	return qb
}

func (qb *Mysql) InsertInto(table string, fields ...string) SqlBuilder {
	qb.Parts = append(qb.Parts, "INSERT INTO", table)
	if len(fields) != 0 {
		fieldsStr := strings.Join(fields, CommaSpace)
		qb.Parts = append(qb.Parts, "(", fieldsStr, ")")
	}
	return qb
}

func (qb *Mysql) Values(vals ...string) SqlBuilder {
	str := strings.Join(vals, CommaSpace)
	qb.Parts = append(qb.Parts, "VALUES", "(", str, ")")
	return qb
}

func (qb *Mysql) Subquery(sub string, alias string) string {
	return fmt.Sprintf("(%s) AS %s", sub, alias)
}

func (qb *Mysql) String() string {
	return strings.Join(qb.Parts, " ")
}
