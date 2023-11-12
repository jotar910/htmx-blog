package storage

import (
	"strings"
	"time"

	"github.com/jotar910/htmx-templ/internal/models"
)

func (mdb *InMemoryDatabase) GetPostsList() *models.ArticleList {
	items := []models.ArticleItem{
		{
			ID:    1,
			Title: "Exploring the Natural Beauty of Yellowstone",
			Image: models.LocalFile{
				Name: "yellowstone.png",
				Url:  "https://images.unsplash.com/photo-1600670942298-b10325b17dea?q=80&w=2574&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
			},
			Summary: "Yellowstone National Park offers a unique blend of natural wonders, from geysers to wildlife. Discover the best trails and spots for an unforgettable adventure.",
			Date:    time.Now(),
		},
		{
			ID:    2,
			Title: "The Culinary Journey Through Italy",
			Image: models.LocalFile{
				Name: "italy-cuisine.png",
				Url:  "https://images.unsplash.com/photo-1590522342323-5d224e217b01?q=80&w=2741&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
			},
			Summary: "Italian cuisine is more than just pizza and pasta. Join us as we dive into the diverse flavors and traditional dishes from various regions of Italy.",
			Date:    time.Now(),
		},
		{
			ID:    3,
			Title: "The Best Kept Secrets of Tokyo City",
			Image: models.LocalFile{
				Name: "tokyo-secrets.png",
				Url:  "https://plus.unsplash.com/premium_photo-1661763066898-6aac7f3758ab?q=80&w=2574&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
			},
			Summary: "Tokyo is a city of endless discovery. We've compiled a list of hidden gems and local favorites that go beyond the typical tourist spots.",
			Date:    time.Now(),
		},
		{
			ID:    4,
			Title: "Hiking the Majestic Trails of Patagonia",
			Image: models.LocalFile{
				Name: "patagonia-trails.png",
				Url:  "https://images.unsplash.com/photo-1546662594-3da3e0603bc9?q=80&w=2670&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
			},
			Summary: "Patagonia's rugged landscapes are a hiker's paradise. Learn about the best times to visit, essential gear, and the most breathtaking trails.",
			Date:    time.Now(),
		},
		{
			ID:    5,
			Title: "A Guide to Sustainable Travel",
			Image: models.LocalFile{
				Name: "sustainable-travel.png",
				Url:  "https://images.unsplash.com/photo-1590598016454-45b7e0ac125c?q=80&w=2343&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
			},
			Summary: "Traveling responsibly is vital for preserving the world's treasures. Discover how you can make a positive impact on your next trip.",
			Date:    time.Now(),
		},
		{
			ID:    6,
			Title: "The Art of French Pastry: A Sweet Journey",
			Image: models.LocalFile{
				Name: "french-pastry.png",
				Url:  "https://images.unsplash.com/photo-1546913686-e667f626c18d?q=80&w=2574&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
			},
			Summary: "French pastries are a feast for the senses. From macarons to croissants, we explore the history and techniques behind France's beloved sweets.",
			Date:    time.Now(),
		},
		{
			ID:    7,
			Title: "Discovering the Ancient Ruins of Machu Picchu",
			Image: models.LocalFile{
				Name: "machu-picchu.png",
				Url:  "https://images.unsplash.com/photo-1590438524133-acf7279afb30?q=80&w=2574&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
			},
			Summary: "Machu Picchu is a testament to the Inca civilization's ingenuity. Join us as we uncover the history and mystery of this ancient wonder.",
			Date:    time.Now(),
		},
		{
			ID:    8,
			Title: "The Vibrant Street Art Scene of Berlin",
			Image: models.LocalFile{
				Name: "berlin-street-art.png",
				Url:  "https://example.com/berlin-street-art.jpg",
			},
			Summary: "Berlin's street art tells the story of the city's cultural and political history. We take a closer look at the murals and the artists behind them.",
			Date:    time.Now(),
		},
		{
			ID:    9,
			Title: "The Ultimate Guide to New York's Coffee Culture",
			Image: models.LocalFile{
				Name: "ny-coffee.png",
				Url:  "https://images.unsplash.com/photo-1611748746228-ddd4c32d966e?q=80&w=2574&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
			},
			Summary: "New York City's coffee scene is as diverse as its boroughs. Find out where to get the best cup and learn about the trends shaping NYC's coffee culture.",
			Date:    time.Now(),
		},
		{
			ID:    10,
			Title: "Chasing the Northern Lights: A Journey to Iceland",
			Image: models.LocalFile{
				Name: "northern-lights.png",
				Url:  "https://images.unsplash.com/photo-1597395529362-361ba4b8ec24?q=80&w=2574&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
			},
			Summary: "The Northern Lights are one of nature's most spectacular displays. Learn the best times and places in Iceland to witness this breathtaking phenomenon.",
			Date:    time.Now(),
		},
		{
			ID:    11,
			Title: "1: The Best United States Destinations to Visit in the Fall",
			Image: models.LocalFile{
				Name: "image.png",
				Url:  "https://101trading.co.uk/wp-content/uploads/2015/04/horizon_00364590-1030x412.jpg",
			},
			Summary: `Are you looking to embrace all those autumnal vibes? Growing up in Southern California, we didn't really get any of those impressive color changes in the leaves nor a serious weather transition.

Are you looking to embrace all those autumnal vibes? Growing up in Southern California, we didn't really get any of those impressive color changes in the leaves nor a serious weather transition. Are you looking to embrace all those autumnal vibes? Growing up in Southern California, we didn't really get any of those impressive color changes in the leaves nor a serious weather transition.

Are you looking to embrace all those autumnal vibes? Growing up in Southern California, we didn't really get any of those impressive color changes in the leaves nor a serious weather transition.`,
			Date: time.Now(),
		},
		{
			ID:    12,
			Title: "2: The Best United States Destinations to Visit in the Fall",
			Image: models.LocalFile{
				Name: "image.png",
				Url:  "https://101trading.co.uk/wp-content/uploads/2015/04/horizon_00364590-1030x412.jpg",
			},
			Summary: `Are you looking to embrace all those autumnal vibes? Growing up in Southern California, we didn't really get any of those impressive color changes in the leaves nor a serious weather transition.

Are you looking to embrace all those autumnal vibes? Growing up in Southern California, we didn't really get any of those impressive color changes in the leaves nor a serious weather transition. Are you looking to embrace all those autumnal vibes? Growing up in Southern California, we didn't really get any of those impressive color changes in the leaves nor a serious weather transition.
Are you looking to embrace all those autumnal vibes? Growing up in Southern California, we didn't really get any of those impressive color changes in the leaves nor a serious weather transition. Are you looking to embrace all those autumnal vibes? Growing up in Southern California, we didn't really get any of those impressive color changes in the leaves nor a serious weather transition.

Are you looking to embrace all those autumnal vibes? Growing up in Southern California, we didn't really get any of those impressive color changes in the leaves nor a serious weather transition.`,
			Date: time.Now(),
		},
		{
			ID:    13,
			Title: "3: The Best United States Destinations to Visit in the Fall",
			Image: models.LocalFile{
				Name: "image.png",
				Url:  "https://101trading.co.uk/wp-content/uploads/2015/04/horizon_00364590-1030x412.jpg",
			},
			Summary: `Are you looking to embrace all those autumnal vibes? Growing up in Southern California, we didn't really get any of those impressive color changes in the leaves nor a serious weather transition.`,
			Date:    time.Now(),
		},
	}
	list := &models.ArticleList{
		Total: len(items),
		Items: items,
	}
	return list
}

func (mdb *InMemoryDatabase) GetPostsListFiltered(
	filters *models.ArticleListFilters,
) *models.ArticleList {
	list := mdb.GetPostsList()

	term := filters.Term
	if term != "" {
		newItems := make([]models.ArticleItem, 0)
		for _, item := range list.Items {
			if strings.Contains(strings.ToLower(item.Title), strings.ToLower(term)) {
				newItems = append(newItems, item)
			}
		}
		list.Items = newItems
		list.Total = len(newItems)
	}

	return list
}
