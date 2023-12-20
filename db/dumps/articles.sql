insert into articles (title, filename, image, summary, timestamp)
values
    ('Exploring the Natural Beauty of Yellowstone',
     'test-1',
     'yellowstone.png',
     'Yellowstone National Park offers a unique blend of natural wonders, from geysers to wildlife. Discover the best trails and spots for an unforgettable adventure.',
     1702146363000),
    ('The Culinary Journey Through Italy',
     'test-2',
     'italy-cuisine.png',
     'Italian cuisine is more than just pizza and pasta. Join us as we dive into the diverse flavors and traditional dishes from various regions of Italy.',
     1702146363000),
    ('The Best Kept Secrets of Tokyo City',
     'test-3',
     'tokyo-secrets.png',
     'Tokyo is a city of endless discovery. We''ve compiled a list of hidden gems and local favorites that go beyond the typical tourist spots.',
     1702146363000),
    ('Hiking the Majestic Trails of Patagonia',
     'test-4',
     'patagonia-trails.png',
     'Patagonia''s rugged landscapes are a hiker''s paradise. Learn about the best times to visit, essential gear, and the most breathtaking trails.',
     1702146363000),
    ('A Guide to Sustainable Travel',
     'test-5',
     'sustainable-travel.png',
     'Traveling responsibly is vital for preserving the world''s treasures. Discover how you can make a positive impact on your next trip.',
     1702146363000),
    ('The Art of French Pastry: A Sweet Journey',
     'test-6',
     'french-pastry.png',
     'French pastries are a feast for the senses. From macarons to croissants, we explore the history and techniques behind France''s beloved sweets.',
     1702146363000),
    ('Discovering the Ancient Ruins of Machu Picchu',
     'test-7',
     'machu-picchu.png',
     'Machu Picchu is a testament to the Inca civilization''s ingenuity. Join us as we uncover the history and mystery of this ancient wonder.',
     1702146363000),
    ('The Vibrant Street Art Scene of Berlin',
     'test-8',
     'berlin-street-art.png',
     'Berlin''s street art tells the story of the city''s cultural and political history. We take a closer look at the murals and the artists behind them.',
     1702146363000),
    ('The Ultimate Guide to New York''s Coffee Culture',
     'test-9',
     'ny-coffee.png',
     'New York City''s coffee scene is as diverse as its boroughs. Find out where to get the best cup and learn about the trends shaping NYC''s coffee culture.',
     1702146363000);

insert into article_tags (article_id, tag_name)
values
    (1, 'Frontend'),
    (2, 'Frontend'),
    (2, 'React'),
    (3, 'Testing'),
    (4, 'Testing'),
    (5, 'Software'),
    (6, 'Review'),
    (7, 'Angular'),
    (7, 'Code Styling'),
    (8, 'Testing'),
    (9, 'Lifestyle'),
    (9, 'Programming');

insert into article_carousel (article_id)
values
    (1),
    (2),
    (3),
    (4),
    (5);

insert into article_highlights(article_id)
values
    (3),
    (4),
    (5);

insert into article_views (article_id, views)
values
    (1, 5),
    (2, 5),
    (3, 0),
    (4, 2),
    (5, 7);
