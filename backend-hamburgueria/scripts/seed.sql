-- Script para inserir dados de exemplo na hamburgueria
-- Baseado nos dados mockados do frontend Vue.js

-- Limpar dados existentes (opcional)
-- DELETE FROM order_items;
-- DELETE FROM orders;
-- DELETE FROM ingredients;
-- DELETE FROM products;
-- DELETE FROM categories;

-- Inserir categorias (baseadas no Menu.vue)
INSERT INTO categories (name, description) VALUES
('Burgers', 'Hambúrgueres tradicionais e especiais'),
('Bebidas', 'Refrigerantes, sucos e outras bebidas'),
('Acompanhamentos', 'Batatas, saladas e outros acompanhamentos'),
('Sobremesas', 'Sobremesas deliciosas'),
('Ingredientes', 'Ingredientes para montagem personalizada');

-- Inserir produtos (baseados no Menu.vue)
INSERT INTO products (name, description, price, category_id, image_url) VALUES
-- Burgers
('Classic Burger', 'Hambúrguer artesanal, queijo cheddar, alface, tomate e molho especial.', 24.90, 1, 'https://images.unsplash.com/photo-1550547660-d9450f859349?auto=format&fit=crop&w=400&q=80'),
('Bacon Deluxe', 'Burger com bacon crocante, queijo e cebola caramelizada.', 29.90, 1, 'https://images.unsplash.com/photo-1519864600265-abb23847ef2c?auto=format&fit=crop&w=400&q=80'),
('Outback Burger', 'Hambúrguer clássico com carne, queijo, alface, tomate e cebola', 25.90, 1, 'https://images.unsplash.com/photo-1550547660-d9450f859349?auto=format&fit=crop&w=400&q=80'),
('Bacon Cheese Burger', 'Hambúrguer com bacon crocante e queijo cheddar', 28.90, 1, 'https://images.unsplash.com/photo-1519864600265-abb23847ef2c?auto=format&fit=crop&w=400&q=80'),
('Veggie Burger', 'Hambúrguer vegetariano com grão-de-bico e legumes', 22.90, 1, 'https://images.unsplash.com/photo-1550547660-d9450f859349?auto=format&fit=crop&w=400&q=80'),

-- Bebidas
('Refrigerante', 'Coca-Cola, Guaraná, Fanta ou Sprite (350ml).', 6.00, 2, 'https://images.unsplash.com/photo-1504674900247-0877df9cc836?auto=format&fit=crop&w=400&q=80'),
('Coca-Cola', 'Refrigerante Coca-Cola 350ml', 6.90, 2, 'https://images.unsplash.com/photo-1504674900247-0877df9cc836?auto=format&fit=crop&w=400&q=80'),
('Suco Natural', 'Suco natural de laranja 300ml', 8.90, 2, 'https://images.unsplash.com/photo-1504674900247-0877df9cc836?auto=format&fit=crop&w=400&q=80'),

-- Acompanhamentos
('Batata Frita', 'Porção individual de batata frita crocante.', 12.90, 3, 'https://images.unsplash.com/photo-1464306076886-debca5e8a6b0?auto=format&fit=crop&w=400&q=80'),
('Onion Rings', 'Anéis de cebola empanados', 14.90, 3, 'https://images.unsplash.com/photo-1464306076886-debca5e8a6b0?auto=format&fit=crop&w=400&q=80'),

-- Sobremesas
('Brownie', 'Brownie de chocolate com sorvete', 15.90, 4, 'https://images.unsplash.com/photo-1464306076886-debca5e8a6b0?auto=format&fit=crop&w=400&q=80'),
('Cheesecake', 'Cheesecake de frutas vermelhas', 18.90, 4, 'https://images.unsplash.com/photo-1464306076886-debca5e8a6b0?auto=format&fit=crop&w=400&q=80');

-- Inserir ingredientes para montagem personalizada (baseados no CustomBurger.vue)
INSERT INTO ingredients (name, price, category) VALUES
-- Pães (baseados no CustomBurger.vue)
('Australiano', 5.00, 'pão'),
('Brioche', 4.00, 'pão'),
('Tradicional', 3.00, 'pão'),

-- Carnes (baseados no CustomBurger.vue)
('Angus 120g', 12.00, 'carne'),
('Frango Grelhado', 10.00, 'carne'),
('Veggie', 9.00, 'carne'),

-- Queijos (baseados no CustomBurger.vue)
('Cheddar', 3.00, 'queijo'),
('Mussarela', 2.50, 'queijo'),
('Sem queijo', 0.00, 'queijo'),

-- Vegetais (baseados no CustomBurger.vue)
('Alface', 1.00, 'vegetais'),
('Tomate', 1.00, 'vegetais'),
('Cebola Crispy', 2.00, 'vegetais'),
('Picles', 1.50, 'vegetais'),

-- Molhos (baseados no CustomBurger.vue)
('Barbecue', 2.00, 'molhos'),
('Maionese da Casa', 1.50, 'molhos'),
('Sem molho', 0.00, 'molhos'),

-- Ingredientes adicionais (do script original)
('Pão Australiano', 0.00, 'pão'),
('Pão Brioche', 2.00, 'pão'),
('Pão Integral', 1.50, 'pão'),
('Carne Angus', 0.00, 'carne'),
('Carne de Cordeiro', 3.00, 'carne'),
('Queijo Gouda', 1.50, 'queijo'),
('Queijo Provolone', 2.00, 'queijo'),
('Queijo Azul', 3.00, 'queijo'),
('Cebola', 0.00, 'vegetais'),
('Pepino', 0.50, 'vegetais'),
('Rúcula', 1.00, 'vegetais'),
('Cogumelos', 2.00, 'vegetais'),
('Maionese', 0.00, 'molhos'),
('Ketchup', 0.00, 'molhos'),
('Mostarda', 0.00, 'molhos'),
('Molho Especial', 1.50, 'molhos'),
('Molho BBQ', 1.00, 'molhos'),
('Molho Ranch', 1.50, 'molhos');

-- Inserir alguns pedidos de exemplo
INSERT INTO orders (customer_name, table_number, total_amount, status, notes) VALUES
('João Silva', 5, 45.80, 'preparing', 'Sem cebola no Classic Burger'),
('Maria Santos', 3, 32.90, 'pending', 'Batata extra crocante'),
('Pedro Costa', 8, 67.70, 'ready', 'Urgente'),
('Ana Oliveira', 2, 28.90, 'delivered', 'Molho especial no lado');

-- Inserir itens dos pedidos de exemplo
INSERT INTO order_items (order_id, product_id, ingredients, quantity, unit_price, total_price, notes) VALUES
(1, 1, '{"observações": "Sem cebola"}', 1, 24.90, 24.90, 'Sem cebola'),
(1, 9, '{}', 1, 12.90, 12.90, ''),
(1, 6, '{}', 1, 6.00, 6.00, ''),
(2, 2, '{}', 1, 29.90, 29.90, ''),
(2, 9, '{"observações": "Extra crocante"}', 1, 12.90, 12.90, 'Batata extra crocante'),
(3, 1, '{}', 2, 24.90, 49.80, ''),
(3, 6, '{}', 2, 6.00, 12.00, ''),
(3, 9, '{}', 1, 12.90, 12.90, ''),
(4, 2, '{}', 1, 29.90, 29.90, 'Molho especial no lado');

-- Atualizar timestamps dos pedidos para parecer mais realista
UPDATE orders SET 
  created_at = NOW() - INTERVAL '2 hours',
  updated_at = NOW() - INTERVAL '1 hour'
WHERE id = 1;

UPDATE orders SET 
  created_at = NOW() - INTERVAL '1 hour',
  updated_at = NOW() - INTERVAL '30 minutes'
WHERE id = 2;

UPDATE orders SET 
  created_at = NOW() - INTERVAL '30 minutes',
  updated_at = NOW() - INTERVAL '10 minutes'
WHERE id = 3;

UPDATE orders SET 
  created_at = NOW() - INTERVAL '2 hours',
  updated_at = NOW() - INTERVAL '1 hour 30 minutes'
WHERE id = 4; 