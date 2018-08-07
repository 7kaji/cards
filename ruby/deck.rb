class Deck
  def initialize
    @cards = %w[♠ ♥ ♦ ♣].product([*1..13]).map do |suit, number|
      Card.new(suit, number)
    end
  end

  def show
    @cards.each {|card| card.show }
  end

  def shuffle!
    @cards.shuffle!
  end

  def save_file
    File.open("deck.txt", "w+") do |f|
      @cards.each {|card| f.puts(card.show) }
    end
  end
end
