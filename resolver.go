package projeto3

import (
	"context"
	"fmt"
	"projeto3/models"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Pedido() PedidoResolver {
	return &pedidoResolver{r}
}

type pedidoResolver struct{ *Resolver }

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Pedidos(ctx context.Context) ([]*models.Pedido, error) {
	lista := []*models.Pedido{
		&models.Pedido{
			ID:        "12",
			ClienteId: 1,
		},
		&models.Pedido{
			ID:        "22",
			ClienteId: 2,
		},
		&models.Pedido{
			ID:        "23",
			ClienteId: 1,
		},
	}
	return lista, nil
}
func (r *queryResolver) Clientes(ctx context.Context) ([]*models.Cliente, error) {
	fmt.Println("Clientes")
	lista := []*models.Cliente{&models.Cliente{ID: "1", Nome: "Jose"},
		&models.Cliente{ID: "2", Nome: "Joao"},
	}
	return lista, nil
}

func (r *pedidoResolver) Cliente(ctx context.Context, obj *models.Pedido) (*models.Cliente, error) {

	fmt.Println("Cliente ID: %d", obj.ID)

	var cliente *models.Cliente

	if obj.ClienteId == 1 {
		cliente = &models.Cliente{ID: "1", Nome: "Jose"}
	} else {
		cliente = &models.Cliente{ID: "2", Nome: "Joao"}
	}

	//lista := &models.Cliente{ID: "1", Nome: "Jose"}

	return cliente, nil
}

func (r *pedidoResolver) Itens(ctx context.Context, obj *models.Pedido) ([]*models.Item, error) {

	fmt.Println("Cliente ID: %d", obj.ID)
	itens := []*models.Item{
		&models.Item{
			Quantidade: 200,
		},
		&models.Item{
			Quantidade: 1200,
		},
		&models.Item{
			Quantidade: 5500,
		},
	}
	/*
		var cliente *models.Cliente

		if obj.ClienteId == 1 {
			cliente = &models.Cliente{ID: "1", Nome: "Jose"}
		} else {
			cliente = &models.Cliente{ID: "2", Nome: "Joao"}
		}

		//lista := &models.Cliente{ID: "1", Nome: "Jose"}
	*/
	return itens, nil
}

func (r *queryResolver) Produtos(ctx context.Context) ([]*models.Produto, error) {

	produtos := []*models.Produto{
		&models.Produto{
			Nome: "sabao",
		},
		&models.Produto{
			Nome: "detergente",
		},
		&models.Produto{
			Nome: "sabonete",
		},
	}
	return produtos, nil

}

func (r *queryResolver) Itens(ctx context.Context) ([]*models.Item, error) {

	itens := []*models.Item{
		&models.Item{
			Valor: 500,
		},
		&models.Item{
			Valor: 600,
		},
	}

	return itens, nil

}
