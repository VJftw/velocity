defmodule ArchitectWeb.V1.BuilderChannel do
  use Phoenix.Channel

  def join("builder:" <> builder_id, %{"token" => token}, socket) do
    case Architect.Builders.authenticate_builder(builder_id, token) do
      {:ok, _builder} ->
        {:ok, socket}

      _ ->
        {:error, %{reason: "unauthorized"}}
    end
  end
end
