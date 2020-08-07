### Comando para criação da tabela produtos

```sql
create table produtos (
	id serial primary key,
	nome varchar,
	descricao varchar,
	preco decimal,
	quantidade integer
)
```

### Comando para popular tabela produtos

```sql
insert into produtos (nome, descricao, preco, quantidade) values ('Camiseta', 'Azul, bem bonita', 39, 5);
insert into produtos (nome, descricao, preco, quantidade) values ('Tênis', 'Confortável', 89, 3);
insert into produtos (nome, descricao, preco, quantidade) values ('Fone', 'Muito bom', 59, 2);
insert into produtos (nome, descricao, preco, quantidade) values ('Produto novo', 'Muito legal', 1.99, 1);
```

### Comando para instalação do postgres driver

```
go get github.com/lib/pq
```