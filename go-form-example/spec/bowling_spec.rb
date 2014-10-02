describe "#score" do
  it "returns 0 for all gutter game" do
    boom = 5
    20.times { boom = boom * 2 }
    boom.should eq(5*2*20)
  end
end
