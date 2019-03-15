from django.urls import path

from api.views import DataPointsPhViewSet, DataPointsPhRetrieveViewSet, DataPointsOxygenViewSet, DataPointsOxygenRetrieveViewSet, DataPointsTemperatureViewSet, DataPointsTemperatureRetrieveViewSet

urlpatterns = [
    path('DataPoints/Ph/', DataPointsPhViewSet.as_view()),
    path('DataPoints/Ph/<int:pk>/', DataPointsPhRetrieveViewSet.as_view()),
    path('DataPoints/Oxygen/', DataPointsOxygenViewSet.as_view()),
    path('DataPoints/Ph/<int:pk>/', DataPointsOxygenRetrieveViewSet.as_view()),
    path('DataPoints/Temperature/', DataPointsTemperatureViewSet.as_view()),
    path('DataPoints/Ph/<int:pk>/', DataPointsTemperatureRetrieveViewSet.as_view()),
]