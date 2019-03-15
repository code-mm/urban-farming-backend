from rest_framework import generics
from rest_framework import mixins

from api.models import DataPointsPh, DataPointsOxygen, DataPointsTemperature
from api.serializers import DataPointsPhSerializer, DataPointsOxygenSerializer, DataPointsTemperatureSerializer

class DataPointsPhViewSet(generics.ListCreateAPIView, generics.CreateAPIView):
    queryset = DataPointsPh.objects.all()
    serializer_class = DataPointsPhSerializer

class DataPointsPhRetrieveViewSet(generics.RetrieveAPIView):
    queryset = DataPointsPh.objects.all()
    serializer_class = DataPointsPhSerializer

class DataPointsOxygenViewSet(generics.ListAPIView, generics.CreateAPIView):
    queryset = DataPointsOxygen.objects.all()
    serializer_class = DataPointsOxygenSerializer

class DataPointsOxygenRetrieveViewSet(generics.RetrieveAPIView):
    queryset = DataPointsOxygen.objects.all()
    serializer_class = DataPointsOxygenSerializer

class DataPointsTemperatureViewSet(generics.ListAPIView, generics.CreateAPIView):
    queryset = DataPointsTemperature.objects.all()
    serializer_class = DataPointsTemperatureSerializer

class DataPointsTemperatureRetrieveViewSet(generics.RetrieveAPIView):
    queryset = DataPointsTemperature.objects.all()
    serializer_class = DataPointsTemperatureSerializer