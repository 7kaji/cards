require 'rspec'
require_relative '../deck'
require_relative '../card'

describe Deck do
  subject { Deck.new }

  context "cards count" do
    it { expect(subject.show.size).to eq 52 }
  end

  context 'first card' do
    it { expect(subject.show.first.show).to eq ' 1 of ♠' }
  end

  context 'last card' do
    it { expect(subject.show.last.show).to eq '13 of ♣' }
  end

  context 'shuffled' do
    before { subject.shuffle! }
    it { expect(subject.show.last.show).not_to eq '13 of ♣' }
    it { expect(subject.show.first.show).not_to eq ' 1 of ♠' }
  end
end
