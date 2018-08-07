require './card'
require './deck'

cards = Deck.new
cards.shuffle!
cards.show.each{|card| puts card.show}
cards.save_file
