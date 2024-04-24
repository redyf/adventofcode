defmodule NotQuiteLispTest do
  use ExUnit.Case
  doctest NotQuiteLisp

  test "santa floor" do
    assert NotQuiteLisp.read_file_and_loop("santa.txt") == {:ok, 280}
  end
end
