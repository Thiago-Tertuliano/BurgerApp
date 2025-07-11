-- Script para adicionar produto padrão para lanches customizados
-- Execute este script no PostgreSQL

-- Verificar se já existe um produto com ID 1
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM products WHERE id = 1) THEN
        -- Inserir produto padrão para lanches customizados
        INSERT INTO products (id, name, description, price, category_id, image_url, is_available) 
        VALUES (1, 'Lanche Personalizado', 'Lanche montado pelo cliente com ingredientes personalizados', 0.00, 1, 'https://images.unsplash.com/photo-1550547660-d9450f859349?auto=format&fit=crop&w=400&q=80', true);
        
        RAISE NOTICE 'Produto padrão para lanches customizados criado com ID 1';
    ELSE
        RAISE NOTICE 'Produto com ID 1 já existe';
    END IF;
END $$;

-- Verificar se existe categoria com ID 1, se não, criar
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM categories WHERE id = 1) THEN
        INSERT INTO categories (id, name, description) 
        VALUES (1, 'Lanches Personalizados', 'Lanches montados pelo cliente');
        RAISE NOTICE 'Categoria padrão criada com ID 1';
    END IF;
END $$;

-- Verificar produtos existentes
SELECT id, name, price, category_id FROM products ORDER BY id; 