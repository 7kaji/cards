class Card
  attr_accessor :suit, :number

  def initialize(suit, number)
    @suit = suit
    @number = number
  end

  def show
    "#{number.to_s.rjust(2)} of #{suit}"
  end
end
