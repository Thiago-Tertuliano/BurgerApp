-- Script para atualizar imagens dos produtos
-- Execute este script no PostgreSQL

-- Ver produtos existentes primeiro
SELECT id, name, description FROM products ORDER BY id;

-- Atualizar imagens com as disponíveis na pasta public:

-- Brownie
UPDATE products SET image_url = '/brownie.jfif' WHERE name ILIKE '%brownie%';

-- Onion Rings  
UPDATE products SET image_url = '/rings.jfif' WHERE name ILIKE '%rings%';

-- Batata Frita
UPDATE products SET image_url = '/fritas.jpg' WHERE name ILIKE '%fritas%' OR name ILIKE '%batata%';

-- Suco
UPDATE products SET image_url = '/suco.jpg' WHERE name ILIKE '%suco%';

-- Refrigerante
UPDATE products SET image_url = '/refri.jpg' WHERE name ILIKE '%refri%';

-- Coca-Cola
UPDATE products SET image_url = '/coca.jpg' WHERE name ILIKE '%coca%';

-- Para todos os outros produtos, usar imagem padrão
UPDATE products SET image_url = 'https://images.unsplash.com/photo-1550547660-d9450f859349?auto=format&fit=crop&w=400&q=80' 
WHERE image_url IS NULL OR image_url = '';

-- Ver resultado
SELECT id, name, image_url FROM products ORDER BY id; 